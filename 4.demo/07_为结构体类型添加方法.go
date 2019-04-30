package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

func main() {
	p := Person{1, "mike", 11}
	p.PrintInfo()

	var p2 Person
	(&p2).SetInfo(2, "jane", 22) // 取地址传参
	p2.PrintInfo()
}

// 带有接收者的函数叫方法
func (tmp Person) PrintInfo() {
	fmt.Println("tmp = ", tmp)
}

// 通过一个函数，给成员赋值
func (p *Person) SetInfo(id int, name string, age int) {
	p.name = name
	p.id = id
	p.age = age
}

// 只要接收者类型不一样，就算函数同名，也是不同的方法，不会冲突
// type long *int  // error 不能直接声明为指针类型，只能在函数使用时声明为指针
type long int

func (tmp *long) test() { // 函数使用时声明为指针

}

type char byte

func (tmp char) test() {

}
