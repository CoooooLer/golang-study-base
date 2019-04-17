package main

import "fmt"

func main() {
	var ch byte // 声明字符类型
	ch = 97
	// 格式输出，%c 以字符方式打印，%d 以整型方式打印
	fmt.Printf("%c, %d\n", ch, ch)

	ch = 'a' // 字符 ， 单引号
	fmt.Printf("%c, %d\n", ch, ch)

	// 大写转小写，小写转大写，大小写相差32
	fmt.Printf("大写: %d, 小写: %d\n", 'A', 'a')
	fmt.Printf("大写转小写: %c\n", 'A'+32)
	fmt.Printf("小写转大写: %c\n", 'a'-32)

	// '\' 以反斜杠开头的字符是转义字符，'\n' 代表换行
	fmt.Printf("hello go%c", '\n')
	fmt.Printf("hello itcast")
}
