package main

import "fmt"

func main() {
	a := 10
	str := "hello"

	// 匿名函数，没有函数名字，函数定义，还没有调用
	f1 := func() {
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}
	f1()

	// 给函数类型起别名
	type FuncType func() // 函数没有参数，没有返回值
	var f2 FuncType
	f2 = f1
	f2()

	// 定义匿名函数同时调用
	func() {
		fmt.Printf("a = %d, str = %s\n", a, str)
	}() // 后面的 （） 代表调用次匿名函数

	// 带参数的匿名函数
	f3 := func(i int, j int) {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}
	f3(11, 22)

	// 定义匿名函数，同时调用
	func(i int, j int) {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}(1, 2)

	// 匿名函数，有参有返回值
	x, y := func(i int, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(123, 321)

	fmt.Printf("x = %d, y = %d\n", x, y)

	// 匿名函数，有参有返回值
	f4 := func(i int, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}
	max, min := f4(111, 222)
	fmt.Printf("max = %d, min = %d\n", max, min)
}
