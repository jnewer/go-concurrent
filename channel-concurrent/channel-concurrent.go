package main

import "fmt"

func worker1(ch chan string) {
	fmt.Println("worker1:doing some work")
	ch <- "worker1:worker done"
}

func worker2(ch chan string) {
	msg := <-ch
	fmt.Println("worker2:received message:", msg)
	fmt.Println("worker2:doing some work")
}
func main() {
	ch := make(chan string)
	go worker1(ch)
	go worker2(ch)
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
