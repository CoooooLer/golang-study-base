package main

import "fmt"

func main() {
	//不同类型变量的声明（定义）
	// var a int
	// var b float32

	// var (
	// 	a int     = 5
	// 	b float64 = 1.2
	// )

	// 自动推导类型
	var (
		a = 3
		b = 3.14
	)

	a, b = 10, 3.14
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// const i int = 10
	// const j float64 = 3.14

	// const (
	// 	i int     = 1
	// 	j float64 = 1.1
	// )

	// 自动推导类型
	const (
		i = 3
		j = 3.14
	)

	fmt.Println("i = ", i)
	fmt.Println("j = ", j)

}
