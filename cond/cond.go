package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	ready := false

	for i := 0; i < 5; i++ {
		go func(i int) {
			mutex.Lock()
			for !ready {
				cond.Wait()
			}

			fmt.Printf("Worker %d is working\n", i)
			mutex.Unlock()
		}(i)
	}

	time.Sleep(time.Second)

	mutex.Lock()
	ready = true
	cond.Broadcast()
	mutex.Unlock()

	time.Sleep(time.Second)
}
