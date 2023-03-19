package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("done")
	}()

	fmt.Println("end")
	time.Sleep(2 * time.Second)
}
