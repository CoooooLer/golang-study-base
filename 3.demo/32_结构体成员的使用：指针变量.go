package main

import "fmt"

type Student struct {
	id   int
	age  int
	name string
	sex  byte
	addr string
}

func main() {
	// 1、指针有合法指向后，才操作成员
	var s Student   // 定义一个普通结构体变量
	var p1 *Student // 定义一个指针变量，保存s的地址
	p1 = &s

	// 通过指针操作成员 p1.id 和 (*p1).id 完全等价，只能使用 . 运算符
	p1.id = 1
	(*p1).name = "goo"
	fmt.Println("p1 = ", p1)

	// 通过 new 申请一个结构体
	p2 := new(Student)
	p2.id = 2
	(*p2).name = "cc"
	fmt.Println("p2 = ", p2)

}
