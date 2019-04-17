package main

import "fmt"

// 多个 defer 会先进后出
func main() {
	defer fmt.Println("aaaaaaaa")
	defer fmt.Println("bbbbbbbbb")
	// 调用一个函数，导致内存出问题
	defer test(0)
	defer fmt.Println("ccccccc")
}

func test(x int) {
	result := 100 / x
	fmt.Println("result = ", result)
}
