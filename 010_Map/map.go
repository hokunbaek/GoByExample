package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)

	m["a"] = 1
	m["b"] = 2
	fmt.Println(m)

	i := m["a"]
	fmt.Println(i)

	fmt.Println(len(m))

	delete(m, "a")
	fmt.Println(m)

	m2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m2)
}
