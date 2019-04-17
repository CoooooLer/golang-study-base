package main

import "fmt"

func main() {
	// 先声明后赋值
	var i int = 1
	i = 10
	i = 20
	i = 30

	fmt.Println("i = ", i)

	// := 自动推导类型
	b := 20
	fmt.Println("b = ", b)

	b = 30
	fmt.Println("b = ", b)

}
