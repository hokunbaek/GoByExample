package main

import "fmt"

func factorialLoop(n int) int {
	if n <= 0 {
		return -1
	}

	result := 1

	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}

func factorialRecursion(n int) int {
	if n <= 0 {
		return -1
	}

	if n == 1 {
		return 1
	}

	return n * factorialRecursion(n-1)
}

func main() {
	fmt.Println(factorialLoop(-5))
	fmt.Println(factorialLoop(5))
	fmt.Println(factorialLoop(10))

	fmt.Println(factorialRecursion(-5))
	fmt.Println(factorialRecursion(5))
	fmt.Println(factorialRecursion(10))
}
