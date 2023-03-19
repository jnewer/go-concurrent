package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()

	select {
	case result := <-ch:
		fmt.Println("操作成功，结果为：", result)
	case <-time.After(2 * time.Second):
		fmt.Println("操作超时，取消任务")
	}
}
