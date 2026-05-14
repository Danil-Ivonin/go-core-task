package main

import (
	"fmt"
	"sync"
)

func mergeChannels(channels ...chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, channel := range channels {
		go func() {
			defer wg.Done()

			for n := range channel {
				out <- n
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	var channels []chan int

	for i := 0; i < 3; i++ {
		channels = append(channels, make(chan int))

		go func() {
			defer close(channels[i])

			for j := 0; j < 100; j++ {
				channels[i] <- j * (i + 1)
			}
		}()
	}

	for res := range mergeChannels(channels...) {
		fmt.Println(res)
	}
}
