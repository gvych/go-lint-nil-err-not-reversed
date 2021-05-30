package analyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:     "gonilerrnotreversed",
	Doc:      "Checks that if err!=nil block not inversed with else block",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func (v *visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	ident, ok := n.(*ast.Ident)
	if ok {
		if ident != nil && ident.Name == "err" {
			v.block = append(v.block, new(bool))
			return v
		}
	}
	return v
}

type visitor struct {
	block []*bool
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.IfStmt)(nil),
	}
	inspector.Preorder(nodeFilter, func(node ast.Node) {
		ifDecl := node.(*ast.IfStmt)

		if res, _ := ifDecl.Cond.(*ast.BinaryExpr); res != nil && res.Op == token.NEQ {
			x, _ := res.X.(ast.Expr).(*ast.Ident)
			y, _ := res.Y.(ast.Expr).(*ast.Ident)
			if (x != nil && x.Name != "err") || (y != nil && y.Name != "nil") { return } //only constructions "if err != nill" accepted

			blockStmt := ifDecl.Body
			if blockStmt != nil {
				v := visitor{}
				ast.Walk(&v, blockStmt)
				if len(v.block) == 0 {
					elseStmt := ifDecl.Else
					if elseStmt != nil {
						ast.Walk(&v, elseStmt)
				    if len(v.block) == 1 {
			    		pass.Reportf(blockStmt.Pos(), "err is not referenced inside error handling block, but referenced in 'else' statement, there is a typo?")
				    }
					}
				}
			}
		}
	})
	return nil, nil
}
