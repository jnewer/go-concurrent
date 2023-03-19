package main

import (
	"fmt"
	"sync"
)

var x = 0
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 100000; i++ {
			mu.Lock()
			x++
			mu.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100000; i++ {
			mu.Lock()
			x--
			mu.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("x = ", x)
}
