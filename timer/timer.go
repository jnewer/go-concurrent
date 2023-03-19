package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	// <-time.After(time.Second)
	// fmt.Println("2 seconds later")

	ticker := time.Tick(time.Second)
	for now := range ticker {
		fmt.Printf("%v\n", now)
	}
}
