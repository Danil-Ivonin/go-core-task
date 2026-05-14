package main

import (
	"slices"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			name:     "basic difference",
			slice1:   []string{"apple", "banana", "cherry", "date"},
			slice2:   []string{"banana", "date"},
			expected: []string{"apple", "cherry"},
		},
		{
			name:     "empty second slice",
			slice1:   []string{"a", "b", "c"},
			slice2:   []string{},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "empty first slice",
			slice1:   []string{},
			slice2:   []string{"a", "b"},
			expected: []string{},
		},
		{
			name:     "no unique elements",
			slice1:   []string{"a", "b"},
			slice2:   []string{"a", "b", "c"},
			expected: []string{},
		},
		{
			name:     "with duplicates",
			slice1:   []string{"apple", "apple", "banana", "cherry"},
			slice2:   []string{"banana"},
			expected: []string{"apple", "apple", "cherry"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := findDiff(tt.slice1, tt.slice2)

			if !slices.Equal(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
