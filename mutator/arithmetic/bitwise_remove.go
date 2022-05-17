package arithmetic

import (
	"go/ast"
	"go/token"
	"go/types"
	"mutesting/mutator"
)

func init() {
	mutator.Register("arithmetic/bitwise_remove", MutatorArithmeticBitwiseRemove)
}

var bitwiseRemoveArithmetic = map[token.Token]bool{
	token.XOR: true,
}

func MutatorArithmeticBitwiseRemove(_ *types.Package, _ *types.Info, node ast.Node) []mutator.Mutation {

	unary, ok := node.(*ast.UnaryExpr)
	if !ok {
		return nil
	}
	original := unary.Op
	X := unary.X
	mutatedX := &ast.UnaryExpr{
		X:  unary.X,
		Op: original,
	}
	if _, ok := bitwiseRemoveArithmetic[original]; !ok {
		return nil
	}
	return []mutator.Mutation{
		{
			Change: func() {
				unary.X = mutatedX
			},
			Reset: func() {
				unary.X = X
			},
		},
	}
}
