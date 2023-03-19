package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	select {
	case ch <- 1:
		fmt.Println("发送成功！")
	default:
		fmt.Println("通道已满，发送失败！")
	}

	select {
	case x := <-ch:
		fmt.Println("接收成功！", x)
	default:
		fmt.Println("通道已空，接受失败！")
	}
}
