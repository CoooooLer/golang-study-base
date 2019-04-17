package main

import "fmt"

// func main() {
// 	var a int

// 	fmt.Printf("请输入一个整数：")

// 	// Scanf 阻塞等待用户的输入
// 	// & 取地址符
// 	// fmt.Scanf("%d", &a)

// 	fmt.Scan(&a)
// 	fmt.Println("a = ", a)
// }

func main() {
	var i float32
	fmt.Printf("请输入一个数：")
	fmt.Scanf("%f", &i)
	fmt.Println("i = ", i)
}
