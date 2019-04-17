package main

import "fmt"

func main() {
	m := map[int]string{1: "go", 2: "c"}
	fmt.Println("m = ", m)

	delete(m, 1) // 删除 key 为 1 的值
	fmt.Println("m = ", m)
}
