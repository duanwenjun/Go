package main

import (
	"fmt"
	"reflect"
)

// 变量标准声明=> var 变量名 变量类型
var name = "张三"
var age = 20
var height int16 = 178

// 常亮
const ProjectType string = "go"

func main() {
	eat("香蕉")
}

// 吃的方法
func eat(name string) {
	fmt.Print("\n")
	fmt.Printf("吃东西 %v ", name)
}

// 全局变量
func globalVariable() {
	fmt.Printf("姓名 %v 年龄 %v 身高 %v", name, age, height)
	fmt.Print("\n")
	fmt.Printf("姓名字段类型 %v 年龄字段 %v 身高字段%v", reflect.TypeOf(name),
		reflect.TypeOf(age), reflect.TypeOf(height))
}

// 常亮
func constMethod() {
	fmt.Println("常量:", ProjectType)
}

// 匿名变量
func anonymousVariable() {
	name, _ := "张三", 18
	_, age := "李四", 20
	fmt.Println(name, age)
}

// 短变量
func shortVariable() {
	// 局部声明变量
	a := 23
	fmt.Println(a)
}
