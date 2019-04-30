package main

import "fmt"

type Humaner interface { // 子集
	Sayhi()
}

type Personer interface { // 超集
	Humaner // 匿名字段继承 Sayhi() 方法
	Sing(lrc string)
}

type Student struct {
	id   int
	name string
}

// Student 实现 Sayhi() 方法
func (tmp *Student) Sayhi() {
	fmt.Printf("Student[%d, %s] sayhi\n", tmp.id, tmp.name)
}

// Student 实现 Sing() 方法
func (tmp *Student) Sing(lrc string) {
	fmt.Printf("Student[%d, %s] Sing %s\n", tmp.id, tmp.name, lrc)
}

func main() {
	var p Personer
	s := Student{123, "gogogogo"}
	p = &s

	p.Sayhi() // 继承过来的方法
	p.Sing("sunshine")

}
