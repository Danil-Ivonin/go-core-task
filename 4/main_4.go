package main

import (
	"fmt"
)

func findDiff(slice1, slice2 []string) []string {
	exists := map[string]struct{}{}

	for _, s := range slice2 {
		exists[s] = struct{}{}
	}

	var res []string
	for _, s := range slice1 {
		if _, ok := exists[s]; !ok {
			res = append(res, s)
		}
	}

	return res
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(findDiff(slice1, slice2))
}
