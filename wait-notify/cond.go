package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	var cv *sync.Cond = sync.NewCond(&mu)

	done := false

	go func() {
		time.Sleep(2 * time.Second)
		mu.Lock()
		done = true
		cv.Signal()
		mu.Unlock()
	}()

	mu.Lock()
	for !done {
		cv.Wait()
	}

	fmt.Println("done")
	mu.Unlock()
}
