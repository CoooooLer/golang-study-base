package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

func main() {
	var p1 Person = Person{1, "go", 11} // 初始化赋值
	fmt.Printf("&p1 = %p\n", &p1)       // %p 打印内存地址

	// p1.setInfoValue(2, "c++", 22)	// 普通传值
	(&p1).setInfoPointer(3, "ggggg", 33) // 取地址传值

	fmt.Println("p1 = ", p1)
}

// 值语义 值传递
func (tmp Person) setInfoValue(id int, name string, age int) {
	fmt.Printf("&tmp = %p\n", &tmp)
	tmp.id = id
	tmp.name = name
	tmp.age = age
	fmt.Println("tmp = ", tmp)
}

// 引用语义 指针传递
func (p *Person) setInfoPointer(id int, name string, age int) {
	fmt.Printf("&p = %p\n", p)
	p.id = id
	p.name = name
	p.age = age
	fmt.Println("p = ", p)
}
