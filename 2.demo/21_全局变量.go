package main

import "fmt"

// a := 10  //error

// 定义在函数外部的变量是全局变量
// 全局变量在任何地方都可以使用
var a int //全局变量只能以 var 的形式声明

func main() {
	a = 11
	fmt.Println("a = ", a)

	test()
}

func test() {
	a = 22
	fmt.Println("a = ", a)
}
