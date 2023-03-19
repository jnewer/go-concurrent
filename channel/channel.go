package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for {
			num, ok := <-ch
			if ok {
				fmt.Println("Received:", num)
			} else {
				fmt.Println("Channel is Closed")
				break
			}
		}
	}()

	fmt.Scanln()
	fmt.Println("Main goroutine is done")
}
