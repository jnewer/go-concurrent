package main

import "fmt"

func main() {
	msgCh := make(chan string)
	go func(ch chan<- string) {
		ch <- "hello from child goroutine"
	}(msgCh)

	msg := <-msgCh
	fmt.Println(msg)
}
