package main

import "fmt"

func main() {
	// a, b, c := MyFunc01()
	// fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	a, b, c := MyFunc03()
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

}

func MyFunc01() (int, int, int) {
	return 1, 2, 3
}

// go 推荐写法
func MyFunc02() (a int, b int, c int) {
	a, b, c = 11, 22, 33
	return
}

func MyFunc03() (a, b, c int) {
	a, b, c = 111, 222, 333
	return
}
