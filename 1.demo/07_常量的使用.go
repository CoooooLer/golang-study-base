package main

import "fmt"

func main() {
	//变量：程序运行期间，可以改变的量
	//常量：程序运行期间，不可以改变的量
	const a int = 10
	//a = 20 //err,常量不允许修改
	fmt.Println("a = ", a)

	const b = 11.2 // 没有使用 :=
	fmt.Printf("b type is %T\n", b)
	fmt.Println("b = ", b)
}
