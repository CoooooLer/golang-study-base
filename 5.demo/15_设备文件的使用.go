package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Stdout.Close()        // 关闭输出
	fmt.Println("are u ok?") //往标准输出设备(屏幕)写内容

	// os.Stdout 标准输出设备
	os.Stdout.WriteString("hello ?")

	// os.Stdin.Close() // 关闭输入
	var a int
	fmt.Println("请输入数字:")
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}
