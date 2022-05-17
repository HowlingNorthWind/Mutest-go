package relational

import (
	"go/ast"
	"go/token"
	"go/types"

	"mutesting/mutator"
)

func init() {
	mutator.Register("relational/negated", MutatorRelationalNegated)
}

var negatedMutations = map[token.Token]token.Token{
	token.GTR: token.LEQ,
	token.LSS: token.GEQ,
	token.GEQ: token.LSS,
	token.LEQ: token.GTR,
	token.EQL: token.NEQ,
	token.NEQ: token.EQL,
}

// MutatorRelationalNegated implements a mutator to improved comparison changes.
func MutatorRelationalNegated(_ *types.Package, _ *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.BinaryExpr)
	if !ok {
		return nil
	}

	original := n.Op
	mutated, ok := negatedMutations[n.Op]
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
