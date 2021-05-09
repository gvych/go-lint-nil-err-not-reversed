package analyzer

import (
	"fmt"
	"go/ast"
	"go/token"
	//"strings"

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

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.IfStmt)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {

		ifDecl := node.(*ast.IfStmt)

		if res := ifDecl.Cond.(*ast.BinaryExpr); res != nil && res.Op == token.NEQ { //}&& res.Y == nil {

			blockStmt := ifDecl.Body

		  ast.Inspect(blockStmt, func(n ast.Node) bool {
		    fmt.Printf("%T\n", n)
				ident, ok := n.(*ast.Ident)
				if ok {
				  if ident != nil && ident.Name == "err" {
  		      pass.Reportf(n.Pos(), "err FINDED in block")
				  }
				}
		    return true
		  })
			return
		}

//
//1 : *ast.IfStmt
//Cond : *ast.BinaryExpr (Op: !=)
//X : *ast.Ident (Name: err)
//Obj : *ast.Object (Kind: var, Name: err)
//Y : *ast.Ident (Name: nil)
//− Body : *ast.BlockStmt
//...
//+ Else : *ast.BlockStmt
//...
//− 0 : *ast.Ident (Name: err)
// Obj : *ast.Object (Kind: var, Name: err)
//...

	})

	return nil, nil
}
