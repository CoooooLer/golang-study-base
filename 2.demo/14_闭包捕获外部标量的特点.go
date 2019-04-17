package main

import "fmt"

func main() {
	a := 1
	str := "hello"
	func() {
		// 闭包以引用的方式捕获外部变量
		a = 777
		str = "go"
		fmt.Printf("内部：a = %d, str = %s\n", a, str)
	}()

	fmt.Printf("外部：a = %d, str = %s\n", a, str)
}
