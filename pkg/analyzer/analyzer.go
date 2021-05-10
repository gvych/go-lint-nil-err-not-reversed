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

type visitor struct{}

func (v visitor) Visit(n ast.Node) ast.Visitor {
		  //ast.Inspect(blockStmt, func(n ast.Node) bool {
		    fmt.Printf("---------Visit enter\n")
		    fmt.Printf("---------%T\n", n)
				ident, ok := n.(*ast.Ident)
				if ok {
				  if ident != nil && ident.Name == "err" {
		        return nil
				  }
				}

		      //pass.Reportf(n.Pos(), "err is not referenced inside block statement")
		    if n == nil {
					fmt.Printf("err is not referenced inside block statement\n")
					return nil
				}
					return v
		  //})
			//return
		//}
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.IfStmt)(nil),
	}
	inspector.Preorder(nodeFilter, func(node ast.Node) {
		ifDecl := node.(*ast.IfStmt)

    fmt.Printf("before BinaryExpr IF\n")
    fmt.Printf("%T\n", ifDecl)
		if res := ifDecl.Cond.(*ast.BinaryExpr); res != nil && res.Op == token.NEQ { //}&& res.Y == nil {
      fmt.Printf("  inside BinaryExpr IF\n")
      fmt.Printf("  %T\n", ifDecl.Body)

			blockStmt := ifDecl.Body
			if blockStmt != nil {
        fmt.Printf("    inside blockStmt NOT nil\n")
        fmt.Printf("    %T\n", blockStmt)

       	nodeIdentFilter := []ast.Node{
       		(*ast.Ident)(nil),
       	}
       	inspector.Preorder(nodeIdentFilter, func(node ast.Node) {
          	ident, ok := node.(*ast.Ident)

       			if ok {
       			  if ident != nil && ident.Name == "err" {
            fmt.Printf("      inside second Preorder\n")
            fmt.Printf("      %T\n", ident)
            fmt.Printf("      %s\n", ident.Name)
           // fmt.Printf("      %d\n", ident.Pos)
		        pass.Reportf(ident.Pos(), "ident.Pos")
       		        return
       			  }
       			}
	      })


      //  v := visitor{}
       // ast.Walk(v, blockStmt)

			}
			return
		}
	})

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


	return nil, nil
}
