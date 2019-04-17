package main

import "fmt"

func main() {
	a := 1
	b := 3.14
	c := 'A'
	d := "hello"

	// %T 操作标量所属类型
	fmt.Printf("%T, %T, %T, %T\n", a, b, c, d)

	fmt.Printf("a = %d, b = %f, c = %c, d = %s\n", a, b, c, d)
	// %v 自动匹配格式输出
	fmt.Printf("a = %v, b = %v, c = %v, d = %v\n", a, b, c, d)

}
