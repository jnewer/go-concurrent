package main

import "fmt"

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Produced:", i)
	}
	close(ch)
}

func consumer(ch chan int) {
	for {
		val, ok := <-ch
		if ok == false {
			break
		}

		fmt.Println("Consumed:", val)
	}
}

func send(ch chan int) {
	fmt.Println("Sending 1st message")
	ch <- 1
	fmt.Println("Sending 2nd message")
	ch <- 2
}
func main() {
	// ch := make(chan int, 2)
	// go producer(ch)
	// consumer(ch)

	ch := make(chan int, 1)
	go send(ch)
	fmt.Println("Receing 1st message")
	fmt.Println(<-ch)
	fmt.Println("Receving 2nd message")
	fmt.Println(<-ch)
}
