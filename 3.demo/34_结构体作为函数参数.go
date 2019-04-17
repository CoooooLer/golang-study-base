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
	s := Student{11, "mike", 33, 'd', "cq"}
	test(s)
	fmt.Println("main: s = ", s)
	test02(&s) // 取地址传参,函数可更改原值
	fmt.Println("指针：s = ", s)
}

// 结构体作为函数参数，只是值传递，函数内部无法更改原值
func test(s Student) {
	s.id = 666
	fmt.Println("test: s = ", s)
}

func test02(p *Student) {
	p.id = 997
}
