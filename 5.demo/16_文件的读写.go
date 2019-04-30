package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	path := "./demo.txt"

	// writeFile(path)
	// readFile(path)
	readFileLine(path)
}

func writeFile(path string) {
	// 打开文件，新建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	// 使用完毕，关闭文件
	defer f.Close()

	var buf string
	for i := 0; i < 10; i++ {
		// buf里面存储的是 i = 0
		buf = fmt.Sprintf("i = %d\n", i)
		f.WriteString(buf)
	}

}

func readFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer f.Close()

	buf := make([]byte, 1024*2) // 创建一个2k的切片

	// n 代表从文件读取内容的长度
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF { // 文件出错，同时没有到结尾
		fmt.Println("err1 = ", err1)
		return
	}

	fmt.Println("buf = ", string(buf[:n])) // 必须转为string类型

}

// 读取一行数据
func readFileLine(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	// 新建一个缓冲区，把内容先放在缓存区
	r := bufio.NewReader(f)

	for {
		// 遇到 \n 结束读取，但是 \n 已经读取到了
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { // 文件已经结束
				break
			}
			fmt.Println("err = ", err)
		}
		fmt.Printf("buf = #%s#\n", string(buf))
	}

}
