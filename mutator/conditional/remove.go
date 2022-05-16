package conditional

import (
	"github.com/avito-tech/go-mutesting/mutator"
	"go/ast"
	"go/token"
	"go/types"
)

func init() {
	mutator.Register("conditional/remove", MutatorConditionalRemove)
}

var removeMutations = map[token.Token]token.Token{
	token.NOT: token.LOR,
	token.LOR: token.LAND,
}

func MutatorConditionalRemove(_ *types.Package, _ *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.UnaryExpr)
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
