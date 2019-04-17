package main

import "fmt"

func main() {
	// 不能转换的类型，叫做不兼容类型
	var flag bool
	flag = true
	fmt.Printf("flag = %t\n", flag)

	// bool类型不能转换为 int
	// fmt.Printf("flag = %d\n", flag)

	// 整型也不能转换为 bool
	// flag = bool(1)
	// fmt.Println("flag = ", flag)

	var ch byte
	ch = 'a' // 字符型本质就是整型
	var t int
	t = int(ch) // 类型转换，把 ch 的值提取出来，转换成 int 再赋值给 t
	fmt.Println("t = ", t)
}
