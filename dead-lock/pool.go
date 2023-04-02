package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	pool := make(chan int, 3)
	for i := 1; i <= 3; i++ {
		pool <- i
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		n := <-pool
		fmt.Println("goroutine 1 acquired resource", n)
		fmt.Println("goroutine 1 is working...")
		time.Sleep(time.Second)
		wg.Done()
		pool <- n
		fmt.Println("goroutine 1 released resource", n)
	}()

	go func() {
		n := <-pool
		fmt.Println("goroutine 2 acquired resource", n)
		fmt.Println("goroutine 2 is working...")
		time.Sleep(time.Second)
		wg.Done()
		pool <- n
		fmt.Println("goroutine 2 released resource", n)
	}()

	go func() {
		n := <-pool
		fmt.Println("goroutine 3 acquired resource", n)
		fmt.Println("goroutine 3 is working...")
		time.Sleep(time.Second)
		wg.Done()
		pool <- n
		fmt.Println("goroutine 3 released resource", n)
	}()

	wg.Wait()
	fmt.Println("All goroutines have finished")
}
