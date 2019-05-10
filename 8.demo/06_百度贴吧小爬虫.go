package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get err: ", err)
		return
	}

	defer resp.Body.Close()

	// 读取网页内容
	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body.Read err: ", err)
			break
		}
		result += string(buf[:n])
	}
	return

}

func DoWork(start, end int) {
	fmt.Printf("爬取%d到%d的页面内容", start, end)

	// http://tieba.baidu.com/f?kw=g-shock&ie=utf-8&pn=50  // 爬取的地址  下一页+50
	for i := start; i <= end; i++ {
		url := "http://tieba.baidu.com/f?kw=g-shock&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		fmt.Println("url = ", url)

		// 爬取数据
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("HttpGet err: ", err)
			continue
		}

		// 把内容写入文件
		fileName := strconv.Itoa(i) + ".html"
		f, err1 := os.Create(fileName)
		if err1 != nil {
			fmt.Println("os.Create err: ", err1)
			continue
		}

		f.WriteString(result)

		f.Close()

	}
}

func main() {
	var start, end int
	fmt.Println("请输入起始页(>=1): ")
	fmt.Scan(&start)
	fmt.Println("请输入结束页(>=起始页): ")
	fmt.Scan(&end)

	DoWork(start, end)

}
