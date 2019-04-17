package main

import "fmt"

func main() {
	fmt.Println("main start")
	funca(5)
	fmt.Println("main end")
}

func funca(a int) {
	fmt.Println("a = ", a)
	if a < 1 {
		return
	}
	funca(a - 1)
}
