package main

import "fmt"

func main() {
	// defer 延迟调用， main 函数结束前调用
	defer fmt.Println("111111111")

	fmt.Println("222222222222")
}
