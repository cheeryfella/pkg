package specification_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cheeryfella/pkg/specification"
)

type TestSpec[t string] struct {
	specification.Spec[string]
	res bool
}

func (s *TestSpec[string]) IsSatisfiedBy(t string) bool {
	_ = t
	return s.res
}

func NewTestSpec(res bool) *TestSpec[string] {
	s := &TestSpec[string]{res: res}
	s.Associate(s)
	return s
}

func TestSpec_Chaining(t *testing.T) {
	tests := map[string]struct {
		build    func() specification.Specification[string]
		expected bool
	}{
		"And then Or — first branch true": {
			build: func() specification.Specification[string] {
				return NewTestSpec(true).And(NewTestSpec(true)).Or(NewTestSpec(false))
			},
			expected: true,
		},
		"And then Or — first branch false, second true": {
			build: func() specification.Specification[string] {
				return NewTestSpec(true).And(NewTestSpec(false)).Or(NewTestSpec(true))
			},
			expected: true,
		},
		"And then Or — both branches false": {
			build: func() specification.Specification[string] {
				return NewTestSpec(true).And(NewTestSpec(false)).Or(NewTestSpec(false))
			},
			expected: false,
		},
		"Or then And — or true, and operand true": {
			build: func() specification.Specification[string] {
				return NewTestSpec(true).Or(NewTestSpec(false)).And(NewTestSpec(true))
			},
			expected: true,
		},
		"Or then And — or true, and operand false": {
			build: func() specification.Specification[string] {
				return NewTestSpec(true).Or(NewTestSpec(false)).And(NewTestSpec(false))
			},
			expected: false,
		},
		"And then Not — negates true result": {
			build: func() specification.Specification[string] {
				return NewTestSpec(true).And(NewTestSpec(true)).Not()
			},
			expected: false,
		},
		"And then Not — negates false result": {
			build: func() specification.Specification[string] {
				return NewTestSpec(true).And(NewTestSpec(false)).Not()
			},
			expected: true,
		},
		"Not then And": {
			build: func() specification.Specification[string] {
				return NewTestSpec(false).Not().And(NewTestSpec(true))
			},
			expected: true,
		},
		"Not then Or — both false after negation": {
			build: func() specification.Specification[string] {
				return NewTestSpec(true).Not().Or(NewTestSpec(false))
			},
			expected: false,
		},
		"NAnd then Or": {
			build: func() specification.Specification[string] {
				return NewTestSpec(true).NAnd(NewTestSpec(true)).Or(NewTestSpec(true))
			},
			expected: true,
		},
		"NOr then Not — double negation": {
			build: func() specification.Specification[string] {
				return NewTestSpec(false).NOr(NewTestSpec(false)).Not()
			},
			expected: false,
		},
		"three-level chain": {
			build: func() specification.Specification[string] {
				return NewTestSpec(true).And(NewTestSpec(true)).Or(NewTestSpec(false)).And(NewTestSpec(true))
			},
			expected: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.build().IsSatisfiedBy("arg"))
		})
	}
}
