package main

import "fmt"

func zeroVal(n int) {
	n = 0
}

func zeroPtr(n *int) {
	*n = 0
}

func main() {
	i := 100
	fmt.Println("init : ", i)

	zeroVal(i)
	fmt.Println("zeroVal : ", i)

	j := 100
	fmt.Println("init : ", j)

	zeroPtr(&j)
	fmt.Println("zeroPtr : ", j)
}
