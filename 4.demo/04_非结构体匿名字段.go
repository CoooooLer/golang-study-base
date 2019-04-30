package main

import "fmt"

type mystr string

type Person struct {
	id   int
	name string
	sex  byte
}

type Student struct {
	Person // 结构体匿名字段
	int    // 基础类型的匿名字段
	mystr
}

func main() {
	var s Student = Student{Person{1, "gooo", 'c'}, 888, "addr"}
	fmt.Printf("s = %+v\n", s)

	s.id = 22
	s.name = "c++"
	s.sex = 'g'
	s.int = 666
	s.mystr = "my"

	fmt.Println(s.id, s.name, s.sex, s.int, s.mystr)

	s.Person = Person{33, "c", 'a'}
	fmt.Println(s.Person, s.int, s.mystr)

}
