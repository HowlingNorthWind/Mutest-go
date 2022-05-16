package conditional

import (
	"mutesting/test"
	"testing"
)

func TestMutatorArithmeticUnary(t *testing.T) {
	test.Mutator(
		t,
		MutatorConditionalReplace,
		"../../testdata/conditional/replace.go",
		2,
	)
}
