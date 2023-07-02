package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			once.Do(func() {
				fmt.Println("Only once")
			})
			fmt.Printf("Hello from goroutine %d\n", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
