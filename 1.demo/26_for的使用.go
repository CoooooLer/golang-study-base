package main

import "fmt"

func main() {
	sum := 1
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)
}
