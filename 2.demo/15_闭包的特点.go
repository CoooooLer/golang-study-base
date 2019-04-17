package main

import "fmt"

func main() {
	// fmt.Println(test01())
	// fmt.Println(test01())
	// fmt.Println(test01())
	// fmt.Println(test01())

	f := test02()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

}

func test01() int {
	// 函数被调用时，i才分配空间，才初始化为 0
	var i int // 没有初始化，默认值为 0
	i++
	return i * i // 函数调用完毕，i 的值才释放
}

func test02() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
