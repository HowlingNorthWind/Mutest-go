package branch

import (
	"go/ast"
	"go/token"
	"go/types"

	"mutesting/mutator"
)

func init() {
	mutator.Register("branch/if_cond", MutatorIfCond)
}

// MutatorIfCond implements a mutator for if and else if branches.
func MutatorIfCond(_ *types.Package, _ *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.IfStmt)
	if !ok {
		return nil
	}
	cond := n.Cond

	trueCond := ast.NewIdent("true")
	falseCond := ast.NewIdent("false")
	trueExpr := &ast.BinaryExpr{
		Op:    token.LOR,
		X:     cond,
		Y:     trueCond,
		OpPos: cond.Pos(),
	}
	falseExpr := &ast.BinaryExpr{
		Op:    token.LAND,
		X:     cond,
		Y:     falseCond,
		OpPos: cond.Pos(),
	}

	return []mutator.Mutation{
		{
			Change: func() {
				n.Cond = trueExpr
			},
			Reset: func() {
				n.Cond = cond
			},
		},
		{
			Change: func() {
				n.Cond = falseExpr
			},
			Reset: func() {
				n.Cond = cond
			},
		},
	}
}
