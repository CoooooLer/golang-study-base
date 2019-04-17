package main

import "fmt"

func main() {
	var i int = 110
	if i == 10 {
		fmt.Println("i 等于 10")
	} else if i > 10 {
		fmt.Println("i 大于 10")
	} else if i < 10 {
		fmt.Println("i 小于 10")
	} else {
		fmt.Println("error")
	}

	if a := 11; a == 11 {
		fmt.Println("a == 11")
	} else {
		fmt.Println("a != 11")
	}

}
