package main

import "fmt"

func main() {
	// 两种方式结果一样，第一种更快
	// 第一种方式
	i := 110
	if i == 10 {
		fmt.Println("i == 10")
	} else if i > 10 {
		fmt.Println("i > 10")
	} else if i < 10 {
		fmt.Println("i < 10")
	} else {
		fmt.Println("error")
	}

	//第二种方式
	if i == 10 {
		fmt.Println("i == 10")
	}

	if i > 10 {
		fmt.Println("i > 10")
	}

	if i < 10 {
		fmt.Println("i < 10")
	}

}
