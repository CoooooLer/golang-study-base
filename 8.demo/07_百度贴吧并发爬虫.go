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
		fmt.Println("htt.Get err: ", err)
		return
	}

	defer resp.Body.Close()

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

func SpiderPage(i int, page chan<- int) { // 通过往管道写入数据控制爬虫
	url := "http://tieba.baidu.com/f?kw=g-shock&ie=utf-8&pn=50" + strconv.Itoa((i-1)*50)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err: ", err)
		// continue
		return
	}

	fileName := strconv.Itoa(i) + ".html"
	f, err1 := os.Create(fileName)
	if err1 != nil {
		fmt.Println("os.Create err: ", err1)
		// continue
		return
	}
	f.WriteString(result)

	page <- i
	f.Close()
}

func DoWork(start, end int) {

	page := make(chan int) //创建一个管道

	for i := start; i <= end; i++ {
		go SpiderPage(i, page) //开子协程

		// 封装成一个函数，便于开协程
		// url := "http://tieba.baidu.com/f?kw=g-shock&ie=utf-8&pn=50" + strconv.Itoa((i-1)*50)
		// result, err := HttpGet(url)
		// if err != nil {
		// 	fmt.Println("HttpGet err: ", err)
		// 	continue
		// }

		// fileName := strconv.Itoa(i) + ".html"
		// f, err1 := os.Create(fileName)
		// if err1 != nil {
		// 	fmt.Println("os.Create err: ", err1)
		// 	continue
		// }
		// f.WriteString(result)

		// f.Close()
		// return
	}

	for i := start; i <= end; i++ {
		fmt.Println("第%d个页面爬取完成！！！", <-page) // 没有数据或者有数据没读取就会阻塞，防止主协程执行完了，子协程被关闭
	}
}

func main() {
	var start, end int
	fmt.Println("请输入起始页:( >=1 )")
	fmt.Scan(&start)
	fmt.Println("请输入结束页:( >=起始页 )")
	fmt.Scan(&end)

	DoWork(start, end)

}
