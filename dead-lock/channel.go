package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		<-ch
	}()

	wg.Wait()
	ch <- 1
	fmt.Println("Data sent to channel")
}
