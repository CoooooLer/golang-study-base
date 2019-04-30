package main

import "fmt"

type Person struct {
	id   int
	name string
	sex  byte
}

// 有个学生，继承Person字段、成员和方法
type Student struct {
	Person
	age  int
	addr string
}

// Person类型，实现了一个方法
func (tmp *Person) PrintInfo() {
	fmt.Printf("id=%d, name=%s, sex=%c\n", tmp.id, tmp.name, tmp.sex)
}

func main() {
	s := Student{Person{1, "mike", 'c'}, 12, "cq"}
	s.PrintInfo()
}
