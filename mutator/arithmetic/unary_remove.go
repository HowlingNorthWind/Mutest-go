package arithmetic

import (
	"go/ast"
	"go/token"
	"go/types"
	"mutesting/mutator"
)

func init() {
	mutator.Register("arithmetic/unary_remove", MutatorArithmeticUnaryRemove)
}

var unaryRemoveArithmetic = map[token.Token]bool{
	token.ADD: true,
	token.SUB: true,
}

func MutatorArithmeticUnaryRemove(_ *types.Package, _ *types.Info, node ast.Node) []mutator.Mutation {

	unary, ok := node.(*ast.UnaryExpr)
	if !ok {
		return nil
	}
	original := unary.Op
	if _, ok := unaryRemoveArithmetic[original]; !ok {
		return nil
	}
	return []mutator.Mutation{
		{
			Change: func() {
				unary.Op = token.ADD
			},
			Reset: func() {
				unary.Op = original
			},
		},
	}
}
