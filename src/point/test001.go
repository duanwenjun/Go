/*
 提供的常用库，有一些常用的方法，方便使用
*/
package main

import "fmt"

type person struct {
	name string
	age  int8
}

//& 取地址操作符 *取值操作符
func newPerson(name string, age int8) *person {
	return &person{
		name: name,
		age:  age,
	}
}

//方法和接收cd
func main() {
	p9 := newPerson("李白", 65)
	fmt.Println("原始值", p9)
	p9.age = 66
	fmt.Println("改变值", p9)

}
