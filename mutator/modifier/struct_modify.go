package modifier

import (
	"go/ast"
	"go/token"
	"go/types"
	"mutesting/mutator"
)

func init() {
	mutator.Register("modifier/struct", MutatorModifierFunction)
}

func MutatorModifierStruct(_ *types.Package, _ *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.GenDecl)
	if !ok {
		return nil
	}
	if !n.Lparen.IsValid() {
		return nil
	}
	if n.Tok != token.TYPE {
		return nil
	}
	//specs := n.Specs
	//for _, spec := range specs {
	//	if typeSpec, ok := spec.(*ast.TypeSpec); ok {
	//
	//	}
	//}
	return nil
}
