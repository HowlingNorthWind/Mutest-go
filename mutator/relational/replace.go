package relational

import (
	"go/ast"
	"go/token"
	"go/types"

	"mutesting/mutator"
)

func init() {
	mutator.Register("relational/replace", MutatorRelationalReplace)
}

var negatedReplaceMutations = map[token.Token]bool{
	token.GTR: true,
	token.LSS: true,
	token.GEQ: true,
	token.LEQ: true,
	token.EQL: true,
	token.NEQ: true,
}

// MutatorRelationalReplace implements a mutator to improved comparison changes.
func MutatorRelationalReplace(_ *types.Package, _ *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.BinaryExpr)
	if !ok {
		return nil
	}

	original := n.Op
	originalX := n.X
	originalY := n.Y
	_, ok = negatedReplaceMutations[n.Op]
	if !ok {
		return nil
	}

	return []mutator.Mutation{
		{
			Change: func() {
				n.X = ast.NewIdent("1")
				n.Y = ast.NewIdent("1")
				n.Op = token.EQL
			},
			Reset: func() {
				n.Op = original
				n.X = originalX
				n.Y = originalY
			},
		},
		{
			Change: func() {
				n.X = ast.NewIdent("1")
				n.Y = ast.NewIdent("1")
				n.Op = token.NEQ
			},
			Reset: func() {
				n.Op = original
				n.X = originalX
				n.Y = originalY
			},
		},
	}
}
