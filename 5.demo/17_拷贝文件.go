package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	list := os.Args // 获取命令行参数
	if len(list) != 3 {
		fmt.Println("usage: xxx srcFile dstFile")
		return
	}

	srcFileName := list[1] // 源文件
	dstFileName := list[2] // 目的文件

	if srcFileName == dstFileName {
		fmt.Println("源文件和目的文件名字不能相同")
		return
	}

	// 只读方式打开源文件
	sF, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	// 新建目的文件
	dF, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}

	// 操作完毕，关闭文件
	defer sF.Close()
	defer dF.Close()

	// 核心处理,从源文件读取内容，往目的文件写，读多少写多少
	buf := make([]byte, 1024*4) // 创建一个4K大小的缓冲区
	for {
		n, err := sF.Read(buf) // 从源文件读取内容
		if err != nil {
			if err == io.EOF { // 内容读取完毕
				break
			}
			fmt.Println("err = ", err)
			return
		}
		// 往目的文件写入，读多少写多少
		dF.Write(buf[:n])
	}

	// 拷贝文件的命令
	// go run 17_拷贝文件.go 12.jpg aaa.jpg

}
