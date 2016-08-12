package main

import "fmt"

func counter(start int) (func() int, func()) {
	c := func() int {
		return start
	}

	i := func() {
		start++
	}

	return c, i
}

func main() {
	c, i := counter(100)
	fmt.Println(c())
	i()
	fmt.Println(c())
	i()
	fmt.Println(c())
}
