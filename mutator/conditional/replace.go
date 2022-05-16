package conditional

import (
	"go/ast"
	"go/token"
	"go/types"
	"mutesting/mutator"
)

func init() {
	mutator.Register("conditional/replace", MutatorConditionalReplace)
}

var replaceMutations = map[token.Token]token.Token{
	token.LAND: token.LOR,
	token.LOR:  token.LAND,
}

func MutatorConditionalReplace(_ *types.Package, _ *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.BinaryExpr)
	if !ok {
		return nil
	}

	original := n.Op
	mutated, ok := replaceMutations[original]
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
