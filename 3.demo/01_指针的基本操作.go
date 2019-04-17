package main

import "fmt"

func main() {
	// 每个变量有两层含义：变量的内存，变量的地址
	var a int = 10
	fmt.Printf("a = %d\n", a)   // 变量的内存
	fmt.Printf("&a = %v\n", &a) // 变量的地址，  %v 自动匹配类型

	// 保存某个变量的地址，需要指针类型， *int 保存int的地址， **int 保存 *int 地址
	// 声明（定义）， 定义就是特殊的声明
	// 定义一个变量 P ，类型为 *int
	var p *int
	p = &a // 指针变量指向谁，就把谁的地址赋值给指针变量
	fmt.Printf("p = %v\n", p)

	*p = 666 //  *p 操作的不是p的内存，是p所指向的内存（就是a）

	fmt.Printf("a = %v\n", a)
	fmt.Printf("*p = %v\n", *p)

}
