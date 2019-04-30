package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

func (p Person) SetInfoValue() {
	fmt.Printf("SetInfoValue: %p %v \n", &p, p)
}

func (p *Person) SetInfoPointer() {
	fmt.Printf("SetPointerValue: %p %v \n", p, p)
}

func main() {
	p := Person{1, "jack", 11}

	p.SetInfoValue() // 传统调用方式

	// 方法值
	pValue := p.SetInfoValue // 调用函数时，无需再传递接收者
	pValue()                 // 等价于 p.SetInfoValue()

	pPointer := p.SetInfoPointer
	pPointer()
}
