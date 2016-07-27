package main

import "fmt"

func main() {
	l := [5]int{1, 2, 3, 4, 5}
	for index, value := range l {
		fmt.Println(index, value)
	}

	fmt.Println()

	s := []int{2, 4, 6, 8, 10}
	for index, value := range s {
		fmt.Println(index, value)
	}

	fmt.Println()

	m := map[string]string{"hello": "world", "foo": "bar", "i": "u"}
	for key, value := range m {
		fmt.Println(key, value)
	}

	fmt.Println()

	str := "HelloWorld"
	for index, value := range str {
		fmt.Printf("%d, %#U\n", index, value)
	}
}
