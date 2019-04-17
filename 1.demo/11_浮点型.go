package main

import "fmt"

func main() {
	// 先声明后赋值
	var i float32
	i = 3.14

	fmt.Println("i = ", i)

	// 自动推导类型
	j := 3.14

	fmt.Printf("j type is %T\n", j)

}
