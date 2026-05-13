package specification_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cheeryfella/pkg/specification"
)

func TestOr_IsSatisfiedBy(t *testing.T) {

	tests := map[string]struct {
		a        specification.Specification[string]
		b        specification.Specification[string]
		expected bool
	}{
		"A & B true": {
			a:        NewTestSpec(true),
			b:        NewTestSpec(true),
			expected: true,
		},
		"A false B true": {
			a:        NewTestSpec(false),
			b:        NewTestSpec(true),
			expected: true,
		},
		"A true B false": {
			a:        NewTestSpec(true),
			b:        NewTestSpec(false),
			expected: true,
		},
		"A & B false": {
			a:        NewTestSpec(false),
			b:        NewTestSpec(false),
			expected: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			underTest := tc.a.Or(tc.b)
			assert.Equal(t, tc.expected, underTest.IsSatisfiedBy("arg"))
		})
	}
}
