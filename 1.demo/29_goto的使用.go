package main

import "fmt"

func main() {
	fmt.Println("11111")

	// goto 可以用在任何地方，但是不能跨函数使用
	goto End // goto是关键字，End 是用户起的名字，他叫标签
	fmt.Println("22222")

End:
	fmt.Printf("33333")
}
