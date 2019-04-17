package main

import "fmt"
import "os"

func main() {
	// 接受用户传递的参数，都是以字符串的形式传递
	list := os.Args
	n := len(list)

	fmt.Println("n = ", n)

	// for i := 0; i < n; i++ {
	// 	fmt.Printf("list[%d] = %s\n", i, list[i])
	// }

	for i, data := range list {
		fmt.Println("list[%d] = %s", i, data)
	}
}
