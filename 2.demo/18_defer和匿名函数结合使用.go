package main

import "fmt"

func main() {
	a := 10
	b := 20
	defer func(a, b int) {
		fmt.Printf("a = %d, b = %d\n", a, b)
	}(a, b) // () 代表调用此匿名函数，把函数传递过去，已经先传递参数，只是没有调用

	a = 111
	b = 222
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func main01() {
	a := 10
	b := 20

	// defer 不管位置在哪里，都只会在函数结束前调用
	defer func() {
		fmt.Printf("a = %d, b = %d\n", a, b)
	}()

	a = 111
	b = 222
	fmt.Printf("外部：a = %d, b = %d\n", a, b)

}
