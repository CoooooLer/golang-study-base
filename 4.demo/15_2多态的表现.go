package main

import "fmt"

type Humaner interface {
	say()
}

type Student struct {
	id   int
	name string
}

func (tmp *Student) say() {
	fmt.Printf("Student[%d, %s] say\n", tmp.id, tmp.name)
}

type Teacher struct {
	age  int
	addr string
}

func (tmp *Teacher) say() {
	fmt.Printf("Teachter[%d, %s] say\n", tmp.age, tmp.addr)
}

type Mystr string

func (tmp *Mystr) say() {
	fmt.Printf("Mystr[%s] say", *tmp)
}

// 定义一个普通函数，函数的参数为接口类型
// 只有一个函数，可以有多种表现，多态
func WhoSay(h Humaner) {
	h.say()
}
func main() {
	s := &Student{123, "jane"}
	t := &Teacher{33, "cq"}
	var str Mystr = "hhhhhh"

	// 调用同一函数，不同表现，多态，多种形态
	WhoSay(s)
	WhoSay(t)
	WhoSay(&str)

	// 创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str

	for _, data := range x {
		data.say()
	}
}
