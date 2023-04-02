package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 1; i < 10; i++ {
		out <- i
		fmt.Printf("producer %d\n", i)
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Printf("consume %d\n", num)
	}
}
func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
	time.Sleep(time.Second)
}
