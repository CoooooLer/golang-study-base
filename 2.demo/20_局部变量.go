package main

import "fmt"

func main() {
	// 定义在 {} 中的变量就是局部变量，只能在 {} 中有效
	// 执行到定义变量的那句代码，才开始分配内存空间，离开作用域自动释放
	// 作用域，变量其作用的范围
	test()

	// a = 11
	// fmt.Println("a = ", a)

	{
		b := 22
		fmt.Println("b = ", b)
	}

	// b = 2222
	// fmt.Println("b = ", b)

	if c := 333; c == 333 {
		fmt.Println("c = ", c)
	}

	// fmt.Println("c = ", c)

}

func test() {
	a := 1
	fmt.Println("a = ", a)
}
