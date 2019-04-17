package calc

import "fmt"

func init() {
	fmt.Println("this is calc init")
}

func Add(a, b int) int {
	return a + b
}

func Minus(a int, b int) int {
	return a - b
}
