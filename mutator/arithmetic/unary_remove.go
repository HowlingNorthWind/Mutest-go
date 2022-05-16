package arithmetic

import (
	"go/ast"
	"go/token"
	"go/types"
	"mutesting/mutator"
)

func init() {
	mutator.Register("arithmetic/unary_remove", MutatorArithmeticUnaryRemove)
}

func MutatorArithmeticUnaryRemove(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {

	unary, ok := node.(*ast.UnaryExpr)
	if !ok {
		return nil
	}
	original := unary.Op
	return []mutator.Mutation{
		{
			Change: func() {
				unary.Op = token.ADD
			},
			Reset: func() {
				unary.Op = original
			},
		},
	}

	//result := make([]mutator.Mutation, 0)
	//switch n := node.(type) {
	//case *ast.BinaryExpr:
	//	if unary, ok := n.X.(*ast.UnaryExpr); ok {
	//		result = append(result, mutator.Mutation{
	//			Change: func() {
	//				n.X = unary.X
	//			},
	//			Reset: func() {
	//				n.X = unary
	//			},
	//		})
	//	}
	//	if unary, ok := n.Y.(*ast.UnaryExpr); ok {
	//		result = append(result, mutator.Mutation{
	//			Change: func() {
	//				n.Y = unary.X
	//			},
	//			Reset: func() {
	//				n.Y = unary
	//			},
	//		})
	//	}
	//case *ast.UnaryExpr:
	//	if unary, ok := n.X.(*ast.UnaryExpr); ok {
	//		result = append(result, mutator.Mutation{
	//			Change: func() {
	//				n.X = unary.X
	//			},
	//			Reset: func() {
	//				n.X = unary
	//			},
	//		})
	//	}
	//case *ast.IndexExpr:
	//	if unary, ok := n.Index.(*ast.UnaryExpr); ok {
	//		result = append(result, mutator.Mutation{
	//			Change: func() {
	//				n.X = unary.X
	//			},
	//			Reset: func() {
	//				n.X = unary
	//			},
	//		})
	//	}
	//case *ast.AssignStmt:
	//	for _, expr := range n.Rhs {
	//		result = append(result, MutatorArithmeticUnaryRemove(pkg, info, expr)...)
	//	}
	//}

	//return result
}
