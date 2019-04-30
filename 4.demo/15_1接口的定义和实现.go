package main

import "fmt"

// 定义接口类型
type Humaner interface {
	say() // 只声明了接口方法，并没有实现，由别的类型(自定义类型)实现
}

type Student struct {
	id   int
	name string
}

// Student 实现 say() 方法
func (tmp *Student) say() {
	fmt.Printf("Student[%d, %s] say \n", tmp.id, tmp.name)
}

type Teacher struct {
	sex  byte
	addr string
}

// Teacher实现 say() 方法
func (tmp *Teacher) say() {
	fmt.Printf("Teacher[%c, %s] say \n", tmp.sex, tmp.addr)
}

type Mystr string

// Mystr 实现 say() 方法
func (tmp Mystr) say() {
	fmt.Printf("Mystr[%s] say\n", tmp)
}

func main() {
	// 定义接口类型的变量
	var h Humaner

	// 只要实现了接口方法的类型，这个类型的变量(接收者类型)就可以给 h 赋值
	s := Student{1, "mike"}
	h = &s
	h.say()

	t := Teacher{'c', "cq"}
	h = &t
	h.say()

	var str Mystr = "me"
	h = str
	h.say()

	// 同一个接口，不同的表现
}
