package arithmetic

import (
	"go/ast"
	"go/token"
	"go/types"
	"mutesting/mutator"
)

func init() {
	mutator.Register("arithmetic/bitwise_insert", MutatorArithmeticBitwiseInsert)
}

var bitwiseInsertArithmetic = map[token.Token]bool{
	token.AND:     true,
	token.OR:      true,
	token.XOR:     true,
	token.AND_NOT: true,
}

func MutatorArithmeticBitwiseInsert(_ *types.Package, _ *types.Info, node ast.Node) []mutator.Mutation {

	n, ok := node.(*ast.BinaryExpr)
	if !ok {
		return nil
	}
	original := n.Op
	if _, ok := bitwiseInsertArithmetic[original]; !ok {
		return nil
	}
	X := n.X
	Y := n.Y
	insertX := &ast.UnaryExpr{
		X:     n.X,
		Op:    token.XOR,
		OpPos: n.X.Pos(),
	}
	insertY := &ast.UnaryExpr{
		X:     n.Y,
		Op:    token.XOR,
		OpPos: n.Y.Pos(),
	}
	var mutations []mutator.Mutation
	if unaryX, ok := X.(*ast.UnaryExpr); !ok || unaryX.Op != token.XOR {
		mutations = append(mutations, mutator.Mutation{
			Change: func() {
				n.X = insertX
			},
			Reset: func() {
				n.X = X
			},
		})
	}
	if unaryY, ok := Y.(*ast.UnaryExpr); !ok || unaryY.Op != token.XOR {
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
