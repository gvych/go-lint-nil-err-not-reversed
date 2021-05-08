package analyzer

import (
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

		//if res := ifDecl.Cond.(*ast.BinaryExpr); res != nil && res.Op != Token.NEQ {
		if res := ifDecl.Cond.(*ast.BinaryExpr); res != nil && res.Op != token.NEQ {
			return
		}
//
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
//
//
//		params := funcDecl.Type.Params.List
//		if len(params) < 2 { // [0] must be format (string), [1] must be args (...interface{})
//			return
//		}
//
//		formatParamType, ok := params[len(params)-2].Type.(*ast.Ident)
//		if !ok { // first param type isn't identificator so it can't be of type "string"
//			return
//		}
//
//		if formatParamType.Name != "string" { // first param (format) type is not string
//			return
//		}
//
//		if formatParamNames := params[len(params)-2].Names; len(formatParamNames) == 0 || formatParamNames[len(formatParamNames)-1].Name != "format" {
//			return
//		}
//
//		argsParamType, ok := params[len(params)-1].Type.(*ast.Ellipsis)
//		if !ok { // args are not ellipsis (...args)
//			return
//		}
//
//		elementType, ok := argsParamType.Elt.(*ast.InterfaceType)
//		if !ok { // args are not of interface type, but we need interface{}
//			return
//		}
//
//		if elementType.Methods != nil && len(elementType.Methods.List) != 0 {
//			return // has >= 1 method in interface, but we need an empty interface "interface{}"
//		}
//
//		if strings.HasSuffix(funcDecl.Name.Name, "f") {
//			return
//		}
//
		pass.Reportf(node.Pos(), "printf-like formatting function '%s' should be named ",
			ifDecl.Pos())
	})

	return nil, nil
}
