package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	// 初始化一个 content
	parent := context.Background()
	//生成一个取消的content
	ctx, cancel := context.WithCancel(parent)
	runTimes := 0
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine Done")
				return
			default:
				fmt.Println("Goroutine Running Times:", runTimes)
				runTimes = runTimes + 1
			}
			if runTimes > 5 {
				fmt.Println("执行超过 5 次,关闭任务")
				cancel()
				wg.Done()
			}
		}
	}(ctx)
	wg.Wait()
}
