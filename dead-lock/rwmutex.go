package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu1,mu2 sync.RWMutex

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		mu1.Lock()
		fmt.Println("goroutine 1 acquired mutex 1")
		mu2.Lock()
		fmt.Println("goroutine 1 acquired mutex 2")
		mu2.Unlock()
		fmt.Println("goroutine 1 released mutex 2")
		mu1.Unlock()
		fmt.Println("goroutine 1 released mutex 1")
		wg.Done()
	}()

	go func() {
		mu2.Lock()
		fmt.Println("goroutine 2 acquired mutex 2")
		mu1.RLock()
		fmt.Println("goroutine 2 acquired mutex 1")
		mu1.RUnlock()
		fmt.Println("goroutine 2 released mutex 1")
		mu2.Unlock()
		fmt.Println("goroutine 2 released mutex 2")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("All goroutines have finished")
}
