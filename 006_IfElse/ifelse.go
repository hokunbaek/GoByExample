package main

import "fmt"

func fizzbuzz(n int) (result string) {
	result = ""

	if n%3 == 0 && n%5 == 0 {
		result = "fizzbuzz"
	} else if n%3 == 0 {
		result = "fizz"
	} else if n%5 == 0 {
		result = "buzz"
	}

	return
}

func main() {
	for i := 1; i < 100; i++ {
		fmt.Printf("%d = %s\n", i, fizzbuzz(i))
	}
}
