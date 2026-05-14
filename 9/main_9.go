package main

import "fmt"

func CubePipeline(input <-chan uint8, output chan<- float64) {
	defer close(output)

	for number := range input {
		value := float64(number)
		output <- value * value * value
	}
}

func main() {
	input := make(chan uint8)
	output := make(chan float64)

	go CubePipeline(input, output)

	go func() {
		defer close(input)
		for _, number := range []uint8{2, 3, 4, 5} {
			input <- number
		}
	}()

	for result := range output {
		fmt.Println(result)
	}
}
