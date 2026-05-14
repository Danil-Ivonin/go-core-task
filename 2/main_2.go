package main

import (
	"fmt"
	"math/rand/v2"
)

func generateSlice(n int) []int {
	res := make([]int, n)
	for i := 0; i < 10; i++ {
		res[i] = rand.IntN(100)
	}
	return res
}

func sliceExample(slice []int) []int {
	var res []int
	for _, n := range slice {
		if n%2 == 0 {
			res = append(res, n)
		}
	}
	return res
}

func addElements(slice []int, num int) []int {
	newSlice := make([]int, len(slice)+1)
	copy(newSlice, slice)
	newSlice[len(slice)] = num
	return newSlice
}

func copySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func removeElement(slice []int, idx int) []int {
	if idx < 0 || idx >= len(slice) {
		panic("index out of range")
	}
	newSlice := make([]int, len(slice)-1)
	copy(newSlice, slice[:idx])
	copy(newSlice[idx:], slice[idx+1:])
	return newSlice
}

func main() {
	originalSlice := generateSlice(10)
	fmt.Printf("originalSlice: %v\n", originalSlice)
	fmt.Printf("sliceExample: %v\n", sliceExample(originalSlice))
	fmt.Printf("addElements: %v\n", addElements(originalSlice, 1))

	copiedSlice := copySlice(originalSlice)
	fmt.Printf("copiedSlice before: %v\n", copiedSlice)
	fmt.Printf("originalSlice before: %v\n", originalSlice)
	copiedSlice[0] = -1
	fmt.Printf("copiedSlice after: %v\n", copiedSlice)
	fmt.Printf("originalSlice after: %v\n", originalSlice)

	fmt.Printf("removeElement 0: %v\n", removeElement(originalSlice, 0))
}
