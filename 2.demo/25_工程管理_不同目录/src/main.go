package main

// 不同目录，包名不一样，调用不同包里面的函数
// 格式： 包名.函数名()
// 函数名必须首字母大写

import (
	"calc" // 只要导入了这个包，就会先执行这个包的 init 函数
	"fmt"
)

// import _ "calc"  // _ 引入包，而不执行包里面的函数，而是为了执行包的 init 函数

func init() {
	fmt.Println("this is main init")
}

func main() {
	res := calc.Add(1, 2)
	fmt.Println("res = ", res)

	fmt.Println("minus = ", calc.Minus(2, 1))
}
