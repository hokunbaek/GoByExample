package main

import (
	"fmt"
	"math"
)

func main() {
	const n = 1000000
	fmt.Println(n)

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(math.Sin(n))
}
