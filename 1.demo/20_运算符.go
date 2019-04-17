package main

import "fmt"

func main() {
	fmt.Println("3 > 2 结果为：", 4 > 3)

	fmt.Println("!(4 != 2) 结果为：", !(4 != 2))

	fmt.Println("true && false 的结果为：", true && false)

	a := 11
	fmt.Println("0 <= a && a <= 100 的结果为：", 0 <= a && a <= 100)
}
