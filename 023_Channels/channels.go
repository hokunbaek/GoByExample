package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		messages <- "Ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}
