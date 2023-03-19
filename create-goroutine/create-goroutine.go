package main

import (
	"fmt"
	"time"
)

func child() {
	for i := 1; i <= 5; i++ {
		fmt.Println("child", i)
		time.Sleep(100 * time.Millisecond)
	}
}
func main() {
	go child()

	for i := 1; i <= 5; i++ {
		fmt.Println("main:", i)
		time.Sleep(100 * time.Millisecond)
	}
}
