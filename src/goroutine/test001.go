package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("action:", i)
}

func main() {
	fmt.Println("cpu number:", runtime.NumCPU())
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
}
