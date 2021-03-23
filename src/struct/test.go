package main

import "fmt"

func main() {
	var p person
	p.age = 10
	p.name = "tom"
	fmt.Println(p)
	//匿名结构体
	var p2 struct {
		color string
		love  string
	}
	p2.color = "颜色"
	p2.love = "喜爱"
	fmt.Printf("p2=%#v\n", p2)
}

type person struct {
	name string
	city string
	age  byte
}
