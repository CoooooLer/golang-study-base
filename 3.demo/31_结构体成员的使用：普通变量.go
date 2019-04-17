package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
	sex  byte
	addr string
}

func main() {
	// 定义一个结构体普通变量
	var s Student

	// 操作成员需要使用 . 运算符
	s.id = 1
	s.name = "go"
	s.age = 22
	s.sex = '1'
	s.addr = "cq"
	fmt.Println("s = ", s)
}
