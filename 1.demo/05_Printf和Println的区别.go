package main

import "fmt"

func main() {
	a := 1
	// 自动换行
	fmt.Println("a = ", a)

	//格式化输出，把a的内容放在%d的位置
	fmt.Printf("a = %d\n", a)

	// := 自动推导类型
	b := 20
	c := 30
	fmt.Println("a = ", a, "b = ", b, "c = ", c)
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

}
