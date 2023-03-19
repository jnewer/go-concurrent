package main

import (
	"fmt"
	"time"
)

func printNum(stopCh <-chan struct{}, n int) {
	for i := 1; i <= n; i++ {
		select {
		case <-stopCh:
			fmt.Println("协程被关闭")
			return
		default:
			fmt.Printf("子协程中的数字：%d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	stopCh := make(chan struct{})
	go printNum(stopCh, 10)

	// 3 秒后关闭协程
	time.Sleep(3 * time.Second)
	close(stopCh)

	// 等待协程结束
	time.Sleep(1 * time.Second)
	fmt.Println("主协程退出")
}
