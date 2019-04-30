package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

type Student struct {
	Person //继承自Person
	sex    byte
	addr   string
	name   string
}

func main() {
	// var s Student = Student{Person{1, "go", 13}, 'd', "cc", "c++"}
	var s1 Student

	// 同名字段就近原则
	// 如果本作用有此成员，就操作此成员
	// 如果无此成员，就找到继承的字段
	s1.id = 1
	s1.name = "go" // 操作的是 Student 里面的 name
	s1.age = 13
	s1.sex = 'c'
	s1.addr = " "

	fmt.Printf("s1 = %+v\n", s1)

	// 显示调用
	s1.Person.name = "c++"
	fmt.Printf("s1 = %+v\n", s1)
}
