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
	s1 := Student{1, "go", 11, 'd', "cq"}
	s2 := Student{1, "go", 11, 'd', "cq"}
	s3 := Student{12, "go", 11, 'd', "cq"}

	fmt.Println("s1 == s2", s1 == s2)
	fmt.Println("s1 == s3", s1 == s3)

	var temp Student
	temp = s3 // 同类型的两个结构体变量可以相互赋值
	fmt.Println("temp = ", temp)
}
