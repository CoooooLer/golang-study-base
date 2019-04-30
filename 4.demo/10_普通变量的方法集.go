package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

func main() {
	p := Person{1, "mike", 12}
	p.setInfoValue()
	p.setInfoPointer() // p自动转化为 (*p) 调用指针方法
}

func (tmp Person) setInfoValue() {
	fmt.Println("setInfoValue")
}

func (tmp *Person) setInfoPointer() {
	fmt.Println("setInfoPointer")
}
