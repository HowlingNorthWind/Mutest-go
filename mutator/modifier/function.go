package modifier

import (
	"go/ast"
	"go/types"
	"mutesting/mutator"
	"strings"
)

var mainFunction = "main"

func init() {
	mutator.Register("modifier/function", MutatorModifierFunction)
}

func MutatorModifierFunction(_ *types.Package, _ *types.Info, node ast.Node) []mutator.Mutation {
	var fun *ast.Ident
	switch n := node.(type) {
	case *ast.FuncDecl:
		fun = n.Name
	case *ast.CallExpr:
		if exp, ok := n.Fun.(*ast.SelectorExpr); ok {
			fun = exp.Sel
		} else if call, ok := n.Fun.(*ast.Ident); ok {
			fun = call
		}
	}

	if fun.Name == mainFunction {
		return nil
	}
	original := fun.Name

	return []mutator.Mutation{
		{
			Change: func() {
				name := fun.Name
				name = strings.ToLower(name[:1]) + name[1:]
				fun.Name = name
			},
			Reset: func() {
				fun.Name = original
			},
		},
	}
}
