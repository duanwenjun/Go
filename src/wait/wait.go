package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(num int) {
			fmt.Println("Test WaitGroup", num)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
