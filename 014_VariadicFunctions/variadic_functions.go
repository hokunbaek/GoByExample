package main

import "fmt"

func sum(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}

	return total
}

func min(n ...int) int {
	if len := len(n); len == 0 {
		return 0
	}

	result := n[0]
	for _, v := range n {
		if v < result {
			result = v
		}
	}

	return result
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(3, 4, 5))

	l := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(l...))

	l2 := []int{3, 4, 5, 4, 2, 9}
	fmt.Println(min(l2...))

	var l3 []int
	fmt.Println(min(l3...))
}
