package main

import "fmt"

var a byte // 全局变量

func main() {
	var a int // 局部变量

	// 1、不同作用域，允许定义同名变量
	// 2、使用变量的原则，就近原则
	fmt.Printf("int：a = %T\n", a) //int
	{
		var a float64
		fmt.Printf("float a = %T\n", a) //float
	}

	test()
}

func test() {
	fmt.Printf("test : a = %T\n", a) //byte
}
