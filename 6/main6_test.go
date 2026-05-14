package main

import "testing"

func TestRandomGenerator(t *testing.T) {
	tests := []struct {
		name  string
		count int
	}{
		{
			name:  "generate 5 numbers",
			count: 5,
		},
		{
			name:  "generate 1 number",
			count: 1,
		},
		{
			name:  "generate 0 numbers",
			count: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ch := randomGenerator(tt.count)

			var numbers []int

			for num := range ch {
				numbers = append(numbers, num)
			}

			if len(numbers) != tt.count {
				t.Errorf("expected %d numbers, got %d", tt.count, len(numbers))
			}
		})
	}
}
