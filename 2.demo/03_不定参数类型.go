package main

import "fmt"

func main() {
	// MyFunc02(1, 2, 3)

	MyFunc03(0, 11)

	// MyFunc04(0, 21)
}

func MyFunc01(a int, b string) { // 固定参数

}

// ...int 不定参数类型， ...type 不定参数类型
func MyFunc02(args ...int) { // 传递的实参可以是0个或多个
	fmt.Println("len(args) = ", len(args))

	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d] = %d\n", i, args[i])
	}

	fmt.Println("=================================")

	for i, data := range args {
		fmt.Printf("args[%d] = %d\n", i, data)
	}
}

// 固定参数一定要传参，不定参数根据需求传参
func MyFunc03(a int, args ...int) {

}

// 不定参数，只能放在形参的最后一个参数
// func MyFunc04(args ...int, a int) { // error

// }
