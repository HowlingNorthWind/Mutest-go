package branch

import (
	"go/ast"
	"go/token"
	"go/types"

	"mutesting/mutator"
)

func init() {
	mutator.Register("branch/insert_fallthrough", MutatorCaseInsertFallthrough)
}

// MutatorCaseInsertFallthrough implements a mutator for case clause
func MutatorCaseInsertFallthrough(_ *types.Package, _ *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.CaseClause)
	if !ok {
		return nil
	}
	list := n.Body

	fallThrough := &ast.BranchStmt{
		Tok: token.FALLTHROUGH,
	}
	mutateList := append(list, fallThrough)

	return []mutator.Mutation{
		{
			Change: func() {
				n.Body = mutateList
			},
			Reset: func() {
				n.Body = list
			},
		},
	}

}
