package main

import "fmt"

func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	i, j := 10, 20
	// 只有 Pritf 可以使用格式化输出， 格式化指带有 %d 的
	fmt.Printf("i = %d, j = %d\n", i, j)

	i, j = j, i
	fmt.Printf("i = %d, j = %d\n", i, j)

	// _ 匿名变量，丢弃数据不处理，_ 匿名变量配合函数返回值使用，才有优势
	var f, d, e int
	f, d, e = test()
	fmt.Printf("c = %d, d = %d, e = %d\n", f, d, e)

	_, d, e = test()
	fmt.Printf("d = %d, e = %d\n", d, e)
}
