package specification_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cheeryfella/pkg/specification"
)

func TestNot_IsSatisfiedBy(t *testing.T) {

	tests := map[string]struct {
		a        specification.Specification[string]
		expected bool
	}{
		"A true": {
			a:        NewTestSpec(true),
			expected: false,
		},
		"A false": {
			a:        NewTestSpec(false),
			expected: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			underTest := tc.a.Not()
			assert.Equal(t, tc.expected, underTest.IsSatisfiedBy("arg"))
		})
	}
}
