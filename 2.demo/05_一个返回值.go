package main

import "fmt"

func main() {
	var a int
	a = MyFunc()
	fmt.Println("a = ", a)

	b := MyFunc()
	fmt.Println("b = ", b)

	c := MyFunc01()
	fmt.Println("c = ", c)

	d := MyFunc02()
	fmt.Println("d = ", d)
}

func MyFunc() int {
	return 6666
}

// 给返回值取一个变量名，go 推荐使用
func MyFunc01() (result int) {
	return 01
}

func MyFunc02() (result int) {
	result = 2
	return
}
