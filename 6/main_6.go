package main

import (
	"fmt"
	"math/rand/v2"
)

func randomGenerator(count int) <-chan int {
	ch := make(chan int) // небуферизированный канал

	go func() {
		defer close(ch)

		for i := 0; i < count; i++ {
			ch <- rand.IntN(100)
		}
	}()

	return ch
}

func main() {
	for num := range randomGenerator(5) {
		fmt.Println(num)
	}
}
