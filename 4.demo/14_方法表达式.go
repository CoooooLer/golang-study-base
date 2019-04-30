package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

func (p Person) SetInfoValue() {
	fmt.Printf("SetInfoValue: %p %v\n", &p, p)
}

func (p *Person) SetInfoPointer() {
	fmt.Printf("SetInfoPointer: %p %v\n", p, p)
}

func main() {
	p := Person{11, "jjjj", 333}

	// 方法值 f := p.SetInfoPointer  隐藏接收者
	// 方法表达式
	f := (Person).SetInfoValue
	f(p) // 显示把接收者传递过去 =====> p.SetInfoValue()

	f2 := (*Person).SetInfoPointer
	f2(&p) // 显示把接收者传递过去 =====> p.SetInfoPointer()
}
