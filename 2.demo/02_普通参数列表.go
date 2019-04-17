package main

import "fmt"

func main() {
	// 调用函数传递的参数叫实参
	MyFunc01(666)
	MyFunc02(1, 2)
	MyFunc03(1, 11)
}

// 定义函数时，函数名后面（）定义的参数叫形参
// 参数传递，只能由实参传递给形参（单向传递），不能反过来
func MyFunc01(a int) {
	//a = 111 // 覆盖传进来的值
	fmt.Println("a = ", a)
}

func MyFunc02(a int, b int) {
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func MyFunc03(a, b int) { // a,b 都是 int 类型
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func MyFunc04(a, b int, c float32, d, e string) {

}

func MyFunc05(a int, b int, c float32, d string, e string) {

}
