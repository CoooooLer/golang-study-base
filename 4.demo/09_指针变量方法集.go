package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

func main() {
	// var p Person = &Person{1, "go", 22}
	// 结构体变量是一个指针变量，它能够调用的方法的集合，简称方法集
	p := &Person{1, "go", 22}
	// 方法自动转换
	p.setInfoPointer()
	(*p).setInfoPointer()

	p.setInfoValue()
	(*p).setInfoValue()
}

func (tmp Person) setInfoValue() {
	fmt.Println("setInfoValue")
}

func (tmp *Person) setInfoPointer() {
	fmt.Println("setInfoPointer")
}
