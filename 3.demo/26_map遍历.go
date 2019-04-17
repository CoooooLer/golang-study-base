package main

import "fmt"

func main() {
	m := map[int]string{1: "name", 2: "age", 3: "sex"}

	// 第一个返回值时 key,第二个返回值是 value,遍历结果是无序的
	for key, value := range m {
		fmt.Printf("%d====>%s\n", key, value)
	}

	// 如何判断一个 key 值是否存在
	// 第一个返回值为 key 所对应的 value
	// 第二个返回值为 key 是否存在的条件，存在 ok 则为 true
	value, ok := m[1]
	if ok == true {
		fmt.Println("m[1] = ", value)
	} else {
		fmt.Println("key 不存在")
	}
}
