package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("wait for message")
		<-ch
		fmt.Println("got message")
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("send message")
	ch <- 1

	time.Sleep(1 * time.Second)
}
