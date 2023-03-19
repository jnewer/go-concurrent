package main

import (
	"fmt"
	"sync"
)

var (
	count int
	m     sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		m.Lock()
		count++
		m.Unlock()
	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go increment(&wg)
	go increment(&wg)

	wg.Wait()
	fmt.Println(count)
}
