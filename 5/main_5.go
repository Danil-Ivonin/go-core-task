package main

import "fmt"

func findIntersections(slice1, slice2 []int) (bool, []int) {
	exists := make(map[int]struct{}, len(slice1))
	for _, v := range slice1 {
		exists[v] = struct{}{}
	}

	var res []int
	for _, v := range slice2 {
		if _, ok := exists[v]; ok {
			res = append(res, v)
		}
	}

	return len(res) != 0, res
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(findIntersections(a, b))
}
