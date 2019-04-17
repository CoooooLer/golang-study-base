package main

import "fmt"

func main() {
	Test(1, 2, 3, 4, 5)
}

func Test(args ...int) {
	// 全部元素传递给MyFunc
	// MyFunc(args...)

	//MyFunc01(args[:2]...) // args[0] ~ args[2] (不包括args[2]) 传递过去
	MyFunc01(args[2:]...) // 从args[2]开始（包括本身），把后面所有元素传递过去
}

func MyFunc(temp ...int) {
	for _, data := range temp {
		fmt.Println("data = ", data)
	}
}

func MyFunc01(temp ...int) {
	for _, data := range temp {
		fmt.Println("data = ", data)
	}
}
