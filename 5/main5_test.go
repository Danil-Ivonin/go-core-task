package main

import (
	"slices"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name           string
		a              []int
		b              []int
		expectedBool   bool
		expectedResult []int
	}{
		{
			name:           "has intersections",
			a:              []int{65, 3, 58, 678, 64},
			b:              []int{64, 2, 3, 43},
			expectedBool:   true,
			expectedResult: []int{64, 3},
		},
		{
			name:           "no intersections",
			a:              []int{1, 2, 3},
			b:              []int{4, 5, 6},
			expectedBool:   false,
			expectedResult: []int{},
		},
		{
			name:           "empty first slice",
			a:              []int{},
			b:              []int{1, 2},
			expectedBool:   false,
			expectedResult: []int{},
		},
		{
			name:           "empty second slice",
			a:              []int{1, 2},
			b:              []int{},
			expectedBool:   false,
			expectedResult: []int{},
		},
		{
			name:           "duplicates in slices",
			a:              []int{1, 2, 2, 3},
			b:              []int{2, 2, 4},
			expectedBool:   true,
			expectedResult: []int{2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ok, result := findIntersections(tt.a, tt.b)

			if ok != tt.expectedBool {
				t.Errorf("expected bool %v, got %v", tt.expectedBool, ok)
			}

			if !slices.Equal(result, tt.expectedResult) {
				t.Errorf("expected result %v, got %v", tt.expectedResult, result)
			}
		})
	}
}
