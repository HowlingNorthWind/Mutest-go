package conditional

import (
	"go/ast"
	"go/token"
	"go/types"
	"mutesting/mutator"
)

func init() {
	mutator.Register("conditional/remove", MutatorConditionalRemove)
}

var removeMutations = map[token.Token]bool{
	token.NOT: true,
}

func MutatorConditionalRemove(_ *types.Package, _ *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.UnaryExpr)
	if !ok {
		return nil
	}

	original := n.Op
	_, ok = removeMutations[original]
	if !ok {
		return nil
	}
	X := n.X
	mutatedX := &ast.UnaryExpr{
		X:  n.X,
		Op: original,
	}
	return []mutator.Mutation{
		{
			Change: func() {
				n.X = mutatedX
			},
			Reset: func() {
				n.X = X
			},
		},
	}
}
