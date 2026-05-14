package main

import (
	"slices"
	"testing"
)

func TestGenerateSlice(t *testing.T) {
	t.Parallel()
	slice := generateSlice(10)

	if len(slice) != 10 {
		t.Errorf("expected length 10, got %d", len(slice))
	}
}

func TestSliceExample(t *testing.T) {
	t.Parallel()

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testSlice := make([]int, len(slice))
	copy(testSlice, slice)

	want := []int{2, 4, 6, 8, 10}
	got := sliceExample(testSlice)

	if !slices.Equal(got, want) {
		t.Errorf("expected %d, got %d", want, got)
	}

	if !slices.Equal(testSlice, slice) {
		t.Errorf("original slice was modified")
	}
}

func TestAddElements(t *testing.T) {
	t.Parallel()

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testSlice := make([]int, len(slice))
	copy(testSlice, slice)

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	got := addElements(testSlice, 11)

	if !slices.Equal(got, want) {
		t.Errorf("expected %d, got %d", want, got)
	}

	if !slices.Equal(testSlice, slice) {
		t.Errorf("original slice was modified")
	}
}

func TestCopySlice(t *testing.T) {
	t.Parallel()

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	copied := copySlice(slice)

	if !slices.Equal(slice, copied) {
		t.Fatalf("want %v, got %v", slice, copied)
	}

	copied[0] = 999
	if slice[0] == 999 {
		t.Fatal("copySlice returned linked slice")
	}
}
func TestRemoveElements(t *testing.T) {
	t.Parallel()
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mid := len(slice) / 2
	tests := []struct {
		name string
		idx  int
		want []int
	}{
		{"removeFirst", 0, []int{2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"removeMid", mid, []int{1, 2, 3, 4, 5, 7, 8, 9, 10}},
		{"removeLast", len(slice) - 1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, tt := range tests {
		testSlice := make([]int, len(slice))
		copy(testSlice, slice)
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := removeElement(testSlice, tt.idx)
			if !slices.Equal(got, tt.want) {
				t.Fatalf("want %v, got %v", tt.want, got)
			}

			if !slices.Equal(testSlice, slice) {
				t.Errorf("original slice was modified")
			}
		})
	}
}

func TestRemoveElementPanic(t *testing.T) {
	t.Parallel()

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	tests := []struct {
		name string
		idx  int
	}{
		{"negative index", -1},
		{"too large index", 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			defer func() {
				r := recover()
				if r == nil {
					t.Fatal("expected panic, got nil")
				}
				if r != "index out of range" {
					t.Fatalf("unexpected panic: %v", r)
				}
			}()
			removeElement(slice, tt.idx)
		})
	}
}
