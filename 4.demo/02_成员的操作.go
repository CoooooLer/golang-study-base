package main

import "fmt"

type Person struct {
	id   int
	name string
	sex  byte
}

type Student struct {
	Person //继承自Person
	age    int
	addr   string
}

func main() {
	var s1 Student = Student{Person{1, "jane", 'c'}, 11, "cq"}
	fmt.Printf("s1 = %+v\n", s1)

	s1.id = 2
	s1.name = "go"
	s1.sex = 'q'
	s1.age = 33
	s1.addr = "cv"
	fmt.Printf("s1 = %+v\n", s1)

	s1.Person = Person{123, "gooo", 'a'}
	fmt.Println(s1.id, s1.name, s1.sex, s1.addr, s1.age)

}
