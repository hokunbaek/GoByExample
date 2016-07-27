package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)
	fmt.Println(len(s))

	c := s[0]
	fmt.Println(c)

	s[0] = "hello"
	fmt.Println(s[0])

	s[1] = "????"
	fmt.Println(s[1])

	s = append(s, "world")
	s = append(s, "!!!!!")
	fmt.Println(s)

	l := s[1:]
	fmt.Println(l)

	l2 := s[:3]
	fmt.Println(l2)
}
