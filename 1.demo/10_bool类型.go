package main

import "fmt"

func main() {
	// 1.声明变量，没有初始化，零值（初始值）为 false
	var a bool
	fmt.Println("a = ", a)

	a = true
	fmt.Println("a = ", a)

	// 2. 自动推导类型
	var b = false
	fmt.Println("b type is %T", b)

	c := true
	fmt.Printf("c type is %T\n", c)
}
