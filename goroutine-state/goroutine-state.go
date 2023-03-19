package main

import (
	"fmt"
	"time"
)

func foo() {
	fmt.Println("New:协程已被创建但还未开始执行任务")
}

func bar() {
	fmt.Println("Runnable:协程正在实行任务")
	time.Sleep(time.Second)
}
func main() {
	go foo()

	go bar()

	ch := make(chan bool)
	go func() {
		fmt.Println("Blocked:协程因为等待channel接收数据而被暂停执行")
		<-ch
	}()

	go func() {
		fmt.Println("Dead:协程执行完成")
	}()

	time.Sleep(2 * time.Second)
}
