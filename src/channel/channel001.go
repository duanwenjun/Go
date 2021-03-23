package main

import "fmt"

func main() {
	// 有缓存通道
	ch2 := make(chan int, 1)
	//把 10 发送channel中
	ch2 <- 10
	//从 channel 中取值赋给 x
	x := <-ch2
	//关闭通道
	close(ch2)
	fmt.Println("通道获取值:", x)
}
