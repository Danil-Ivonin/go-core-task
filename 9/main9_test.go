package main

import (
	"slices"
	"testing"
	"time"
)

func TestCubePipelineProcessesNumbers(t *testing.T) {
	t.Parallel()

	input := make(chan uint8)
	output := make(chan float64)

	go CubePipeline(input, output)

	go func() {
		defer close(input)
		for _, number := range []uint8{2, 3, 4, 5} {
			input <- number
		}
	}()

	var got []float64
	for value := range output {
		got = append(got, value)
	}

	want := []float64{8, 27, 64, 125}
	if !slices.Equal(got, want) {
		t.Fatalf("CubePipeline() = %v, want %v", got, want)
	}
}

func TestCubePipelineClosesOutputForEmptyInput(t *testing.T) {
	t.Parallel()

	input := make(chan uint8)
	output := make(chan float64)

	go CubePipeline(input, output)
	close(input)

	select {
	case value, ok := <-output:
		if ok {
			t.Fatalf("output channel is open, received %v", value)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("output channel was not closed")
	}
}

func TestCubePipelineEmitsValueWhenInputArrives(t *testing.T) {
	t.Parallel()

	input := make(chan uint8)
	output := make(chan float64)

	go CubePipeline(input, output)

	select {
	case value := <-output:
		t.Fatalf("received %v before input was sent", value)
	case <-time.After(20 * time.Millisecond):
	}

	input <- 6

	select {
	case got := <-output:
		if got != 216 {
			t.Fatalf("CubePipeline() emitted %v, want 216", got)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("CubePipeline() did not emit a value after input arrived")
	}

	close(input)
	for range output {
	}
}
