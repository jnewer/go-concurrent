package main

import (
	"fmt"
	"sync"
)

var count int
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			count++
			wg.Done()
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
