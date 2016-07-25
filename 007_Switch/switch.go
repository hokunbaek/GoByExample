package main

import "fmt"

func fizzbuzz(n int) (result string) {
	switch {
	case n%3 == 0 && n%5 == 0:
		result = "fizzbuzz"
	case n%3 == 0:
		result = "fizz"
	case n%5 == 0:
		result = "buzz"
	default:
		result = ""
	}

	return
}

func main() {
	for i := 1; i < 100; i++ {
		fmt.Printf("i = %d, result = %s\n", i, fizzbuzz(i))
	}
}
