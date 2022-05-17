package conditional

import (
	"go/ast"
	"go/token"
	"go/types"
	"mutesting/mutator"
)

func init() {
	mutator.Register("conditional/insert", MutatorConditionalInsert)
}

var conditionalInsert = map[token.Token]bool{
	token.LAND: true,
	token.LOR:  true,
}

func MutatorConditionalInsert(_ *types.Package, _ *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.BinaryExpr)
	if !ok {
		return nil
	}
	original := n.Op
	if _, ok := conditionalInsert[original]; !ok {
		return nil
	}
	X := n.X
	Y := n.Y
	insertX := &ast.UnaryExpr{
		X:     n.X,
		Op:    token.NOT,
		OpPos: n.X.Pos(),
	}
	insertY := &ast.UnaryExpr{
		X:     n.Y,
		Op:    token.NOT,
		OpPos: n.Y.Pos(),
	}
	var mutations []mutator.Mutation
	if unaryX, ok := X.(*ast.UnaryExpr); !ok || unaryX.Op != token.NOT {
		mutations = append(mutations, mutator.Mutation{
			Change: func() {
				n.X = insertX
			},
			Reset: func() {
				n.X = X
			},
		})
	}
	if unaryY, ok := Y.(*ast.UnaryExpr); !ok || unaryY.Op != token.NOT {
		mutations = append(mutations, mutator.Mutation{
			Change: func() {
				n.Y = insertY
			},
			Reset: func() {
				n.Y = Y
			},
		})
	}
	return mutations
}
