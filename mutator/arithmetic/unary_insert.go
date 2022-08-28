package arithmetic

import (
	"go/ast"
	"go/token"
	"go/types"
	"mutesting/mutator"
)

func init() {
	mutator.Register("arithmetic/base_unary_insert", MutatorArithmeticUnaryInsert)
}

var unaryInsertArithmetic = map[token.Token]bool{
	token.ADD: true,
	token.SUB: true,
	token.MUL: true,
	token.QUO: true,
	token.REM: true,
}

func MutatorArithmeticUnaryInsert(_ *types.Package, _ *types.Info, node ast.Node) []mutator.Mutation {

	n, ok := node.(*ast.BinaryExpr)
	if !ok {
		return nil
	}
	original := n.Op
	if _, ok := unaryInsertArithmetic[original]; !ok {
		return nil
	}
	X := n.X
	Y := n.Y
	insertX := &ast.UnaryExpr{
		X:     n.X,
		OpPos: n.X.Pos(),
	}
	insertY := &ast.UnaryExpr{
		X:     n.Y,
		OpPos: n.Y.Pos(),
	}
	var mutations []mutator.Mutation
	if unaryX, ok := X.(*ast.UnaryExpr); !ok || (unaryX.Op != token.ADD && unaryX.Op != token.SUB) {
		mutations = append(mutations, []mutator.Mutation{
			{
				Change: func() {
					insertX.Op = token.ADD
					n.X = insertX
				},
				Reset: func() {
					n.X = X
				},
			},
			{
				Change: func() {
					insertX.Op = token.SUB
					n.X = insertX
				},
				Reset: func() {
					n.X = X
				},
			},
		}...)
	}
	if unaryY, ok := Y.(*ast.UnaryExpr); !ok || (unaryY.Op != token.ADD && unaryY.Op != token.SUB) {
		mutations = append(mutations, []mutator.Mutation{
			{
				Change: func() {
					insertY.Op = token.ADD
					n.Y = insertY
				},
				Reset: func() {
					n.Y = Y
				},
			},
			{
				Change: func() {
					insertY.Op = token.SUB
					n.Y = insertY
				},
				Reset: func() {
					n.Y = Y
				},
			},
		}...)
	}
	return mutations
}
