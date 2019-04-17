package main

import "fmt"

func main() {
	m := map[int]string{1: "ggoo", 2: "ccccc"}
	fmt.Println("m = ", m)
	test(m)
	fmt.Println("m = ", m)
}

// 传递的是原map的值，删除的也是原map的值
func test(m map[int]string) {
	delete(m, 1)
}
