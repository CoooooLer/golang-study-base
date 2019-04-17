package main

import "fmt"

func main() {
	//函数的调用： 函数名（）
	MyFunc()
}

// 无参无返回值函数的定义
func MyFunc() {
	a := 0111
	fmt.Println("a = ", a)
}
