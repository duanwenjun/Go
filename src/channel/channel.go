package main

import "fmt"

func main() {
	i, i2 := calc(1, 2)
	fmt.Println("求和", i)
	fmt.Println("求减", i2)
}
func intSum(x ...int) int {
	fmt.Println(x)
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}
