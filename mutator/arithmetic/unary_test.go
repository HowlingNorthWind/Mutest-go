package arithmetic

import (
	"testing"

	"mutesting/test"
)

func TestMutatorArithmeticUnary(t *testing.T) {
	test.Mutator(
		t,
		MutatorArithmeticUnary,
		"../../testdata/arithmetic/unary.go",
		2,
	)
}
