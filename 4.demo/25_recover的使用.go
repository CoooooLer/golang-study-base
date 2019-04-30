package main

import "fmt"

func main() {
	test(11)
}

func test(i int) {
	defer func() {
		// recover() ，产生了 panic 程序不会崩溃
		// fmt.Println(recover())  // 通过 recover() 打印错误信息
		if err := recover(); err != nil {
			fmt.Println("err:")
			fmt.Println(err)
		}
	}() // () 调用匿名函数自身

	var a [10]int
	a[i] = 123 // 当 i 大于10就会越界，产生一个 panic ,导致程序崩溃
}
