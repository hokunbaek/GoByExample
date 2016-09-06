package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 1")
	}
}