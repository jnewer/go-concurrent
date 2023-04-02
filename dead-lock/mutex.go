package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu1, mu2 sync.Mutex
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		mu1.Lock()
		fmt.Println("groutine 1 acquired mutex 1")
		mu2.Lock()
		fmt.Println("groutine 1 acquired mutex 2")
		mu2.Unlock()
		fmt.Println("groutine 1 released mutex 2")
		mu1.Unlock()
		fmt.Println("groutine 1 released mutex 1")
		wg.Done()
	}()

	go func() {
		mu2.Lock()
		fmt.Println("groutine 1 acquired mutex 2")
		mu1.Lock()
		fmt.Println("groutine 1 acquired mutex 1")
		mu1.Unlock()
		fmt.Println("groutine 1 released mutex 1")
		mu2.Unlock()
		fmt.Println("groutine 1 released mutex 2")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("All goroutines have finished")
}
