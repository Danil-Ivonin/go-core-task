package main

import (
	"slices"
	"testing"
)

func createChannel(values []int) chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, value := range values {
			ch <- value
		}
	}()

	return ch
}

func TestMergeChannels(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			name: "multiple channels",
			input: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "single channel",
			input: [][]int{
				{10, 20, 30},
			},
			expected: []int{10, 20, 30},
		},
		{
			name: "empty channels",
			input: [][]int{
				{},
				{},
			},
			expected: []int{},
		},
		{
			name: "mixed empty and non-empty channels",
			input: [][]int{
				{},
				{1},
				{},
				{2, 3},
			},
			expected: []int{1, 2, 3},
		},
		{
			name:     "no channels",
			input:    [][]int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var channels []chan int

			for _, values := range tt.input {
				channels = append(channels, createChannel(values))
			}

			var result []int

			for value := range mergeChannels(channels...) {
				result = append(result, value)
			}

			slices.Sort(result)

			if !slices.Equal(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
