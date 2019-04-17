package main

import "fmt"

func main() {
	var sum int
	// sum = test(1)
	sum = test01(100)
	fmt.Println("sum = ", sum)

}

func test(i int) (sum int) {
	if i == 100 {
		return 100
	}
	sum += i + test(i+1)
	return
}

func test01(i int) int {
	if i == 1 {
		return 1
	}
	return i + test01(i-1)
}
