package main

import "fmt"

func main() {
	m1 := map[int]string{1: "go", 2: "c++"}
	fmt.Println("m1 = ", m1)

	// 赋值操作，如果 key 值已经存在，则修改内容
	m1[1] = "gooo"
	fmt.Println("m1 = ", m1)

	m1[5] = "ccccccccc" // 追加，map底层自动扩容，和append类似
	fmt.Println("m1 = ", m1)
}
