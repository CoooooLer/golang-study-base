package main

import "fmt"

func main() {
	var result int
	result = Add(1, 2)
	fmt.Println("result = ", result) // 传统调用方式

	var fTest FuncType     // 声明一个函数类型的变量，变量名叫 fTest
	fTest = Add            // 是变量就可以赋值
	result = fTest(11, 22) // 等价于 Add(11, 22)
	fmt.Println("resultF = ", result)

	fTest = Minus
	result = fTest(22, 11) // 等价于 Minus(22, 11)
	fmt.Println("result = ", result)
}

// 函数也是一种数据类型，通过 type 给一个函数类型起名
// FuncType 它是一个函数类型
type FuncType func(int, int) int // 没有函数名字，没有 {}  ，类似指针

func Add(a int, b int) int {
	return a + b
}

func Minus(a int, b int) int {
	return a - b
}
