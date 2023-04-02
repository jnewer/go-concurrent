package main

import (
	"fmt"
	"sync"
)

var count int
var rwmutex sync.RWMutex
var wg sync.WaitGroup

func read(i int) {
	defer wg.Done()
	rwmutex.RLock()
	defer rwmutex.RUnlock()
	fmt.Printf("read %v: count=%v\n", i, count)
}

func write(i int) {
	defer wg.Done()
	rwmutex.Lock()
	defer rwmutex.Unlock()
	count = i
	fmt.Printf("write %v:count=%v\n", i, count)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read(i)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go write(i)
	}

	wg.Wait()
}
