package arithmetic

import (
	"go/ast"
	"go/token"
	"go/types"
	"mutesting/mutator"
)

func init() {
	mutator.Register("arithmetic/base_unary", MutatorArithmeticUnary)
}

var unaryMutations = map[token.Token]token.Token{
	token.ADD: token.SUB,
	token.SUB: token.ADD,
}

func MutatorArithmeticUnary(_ *types.Package, _ *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.UnaryExpr)
	if !ok {
		return nil
	}

	original := n.Op
	mutated, ok := unaryMutations[original]
	if !ok {
		return nil
	}

	return []mutator.Mutation{
		{
			Change: func() {
				n.Op = mutated
			},
			Reset: func() {
				n.Op = original
			},
		},
	}
}
