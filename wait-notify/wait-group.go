package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("groutine 1")
		time.Sleep(1 * time.Second)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("groutine 2")
		time.Sleep(2 * time.Second)
	}()

	wg.Wait()
	fmt.Println("all goroutine done")

}
