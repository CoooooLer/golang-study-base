package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
	sex  byte
	addr string
}

func main() {
	// 顺序初始化，每个成语必须赋值
	var s1 Student = Student{1, "mike", 14, 'm', "地球"}
	fmt.Println("s1 = ", s1)

	// 指定成员初始化，没有初始化的成员，自动赋值为0
	s2 := Student{name: "go", sex: 's'}
	fmt.Println("s2 = ", s2)
}
