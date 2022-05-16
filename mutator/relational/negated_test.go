package relational

import (
	"testing"

	"mutesting/test"
)

func TestMutatorConditionalNegated(t *testing.T) {
	test.Mutator(
		t,
		MutatorConditionalNegated,
		"../../testdata/relational/negated.go",
		6,
	)
}
