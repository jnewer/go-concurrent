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

	for n := range ch {
		fmt.Println(n)
	}
}
