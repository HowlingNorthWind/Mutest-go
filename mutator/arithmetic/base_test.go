package arithmetic

import (
	"testing"

	"mutesting/test"
)

func TestMutatorArithmeticBase(t *testing.T) {
	test.Mutator(
		t,
		MutatorArithmeticBase,
		"../../testdata/arithmetic/base.go",
		5,
	)
}
