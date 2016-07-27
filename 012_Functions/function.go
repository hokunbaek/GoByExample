package main

import "fmt"

func plus(n1 int, n2 int) int {
	return n1 + n2
}

func plusplus(n1, n2, n3 int) int {
	return n1 + n2 + n3
}

func main() {
	fmt.Println(plus(10, 20))
	fmt.Println(plusplus(10, 20, 30))
}
