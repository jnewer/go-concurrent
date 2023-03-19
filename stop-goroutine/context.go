package main

import (
	"context"
	"fmt"
	"time"
)

func printNum(ctx context.Context, n int) {
	for i := 1; i <= n; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("协程被取消")
			return
		default:
			fmt.Printf("子协程中的数字：%d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go printNum(ctx, 10)

	// 3 秒后取消协程
	time.Sleep(3 * time.Second)
	cancel()

	// 等待协程结束
	time.Sleep(1 * time.Second)
	fmt.Println("主协程退出")
}
