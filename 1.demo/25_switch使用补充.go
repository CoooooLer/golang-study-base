package main

import "fmt"

func main() {
	switch i := 6; i { // switch 后面可以直接声明变量，并作为条件变量
	case 1:
		fmt.Println("1")
	case 2, 3, 4:
		fmt.Println("2, 3, 4")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("more")
	}

	scope := 0
	switch { // switch 后面可以不写变量
	case scope > 90: // case 后面可以放条件
		fmt.Println("> 90")
	case scope > 60:
		fmt.Println("> 60")
	case scope < 60:
		fmt.Println("< 60")
	default:
		fmt.Println("default")

	}
}
