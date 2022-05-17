package branch

import (
	"go/ast"
	"go/token"
	"go/types"
	"mutesting/mutator"
)

func init() {
	mutator.Register("branch/remove_fallthrough", MutatorCaseRemoveFallthrough)
}

// MutatorCaseRemoveFallthrough implements a mutator for case clause
func MutatorCaseRemoveFallthrough(_ *types.Package, _ *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.CaseClause)
	if !ok {
		return nil
	}
	list := n.Body

	lastIdx := len(list) - 1
	if lastIdx < 0 {
		return nil
	}
	fallThrough, ok := list[lastIdx].(*ast.BranchStmt)
	if !ok {
		return nil
	}
	if fallThrough.Tok != token.FALLTHROUGH {
		return nil
	}
	mutateList := list[:lastIdx]
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
