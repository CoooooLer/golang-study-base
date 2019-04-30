package main

import "fmt"

type Person struct {
	id   int
	name string
	sex  byte
}

type Student struct {
	Person //继承 Person 的属性
	age    int
	addr   string
}

// Person 类型实现了一个 PrintInfo() 方法
func (tmp *Person) PrintInfo() {
	fmt.Println("Person tmp = ", tmp)
}

// Student 类型也实现了一个 PrintInfo() 方法
func (tmp *Student) PrintInfo() {
	fmt.Println("Student tmp = ", tmp)
}

func main() {
	s := Student{Person{1, "jane", 'd'}, 111, "cc"}
	// 就近原则：先找最近的方法，如果没有就调用继承的方法
	s.PrintInfo()

	// 显示调用继承的方法
	s.Person.PrintInfo()
}
