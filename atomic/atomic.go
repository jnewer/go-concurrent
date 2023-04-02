package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var count int32 = 0
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000000; j++ {
				atomic.AddInt32(&count, 1)
			}
		}()
	}

	time.Sleep(time.Second)

	fmt.Println(count)
}
