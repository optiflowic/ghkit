package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Contains(t *testing.T) {
	type testCase[T comparable] struct {
		name     string
		list     []T
		target   T
		expected bool
	}

	t.Run("int", func(t *testing.T) {
		tests := []testCase[int]{
			{"found", []int{1, 2, 3}, 2, true},
			{"not found", []int{1, 2, 3}, 4, false},
			{"empty list", []int{}, 1, false},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := Contains(tt.list, tt.target)
				assert.Equal(t, tt.expected, got)
			})
		}
	})

	t.Run("string", func(t *testing.T) {
		tests := []testCase[string]{
			{"found", []string{"a", "b", "c"}, "b", true},
			{"not found", []string{"a", "b", "c"}, "x", false},
			{"empty list", []string{}, "a", false},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := Contains(tt.list, tt.target)
				assert.Equal(t, tt.expected, got)
			})
		}
	})
}
