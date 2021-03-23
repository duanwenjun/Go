package main

import (
	"fmt"
	"time"
)

func action(i int) {
	fmt.Println("开始执行", i)
	fmt.Println("Test Goroutine", i)
}
func main() {
	go action(0)
	fmt.Println("执行完成")
	time.Sleep(1 * time.Second)

}
