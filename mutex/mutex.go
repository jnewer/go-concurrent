package main

import (
	"fmt"
	"sync"
)

var count int
var mutex sync.Mutex

func increament() {
	mutex.Lock()
	count++
	mutex.Unlock()
}

func main() {
	for i := 0; i < 10; i++ {
		go increament()
	}

	fmt.Printf("count:%d\n", count)
}
