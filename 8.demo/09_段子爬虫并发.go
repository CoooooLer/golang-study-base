package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	buf := make([]byte, 10245*4)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf[:n])
	}
	return
}

// 获取每一个笑话的title和content
func SpiderOneJoy(url string) (title, content string, err error) {
	result, err1 := HttpGet(url)
	if err1 != nil {
		err = err1
		return
	}

	// 获取h1中的title
	// <h1>会放坏															</h1>
	re1 := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if re1 == nil {
		fmt.Println("regexp.MustCompile title err ")
		return
	}

	// title
	tmpTitle := re1.FindAllStringSubmatch(result, 1) // 因为页面上有两个h1标签，只取第一个
	for _, data := range tmpTitle {
		title = data[1]
		title = strings.Replace(title, "\t", "", -1)
		// title = strings.Replace(title, " ", "", -1)
		title = strings.Replace(title, "\b", "", -1)
		title = strings.Replace(title, "\r", "", -1)
	}

	// 获取content
	// <div class="content-txt pt10">女：“我购物车里的那些水果牛奶饮料零食你赶紧给我买！” 男：“着什么急呀？” 女：“天越来越热了，会放坏的。”								<a id="prev"
	re2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev"`)
	if re2 == nil {
		fmt.Println("regexp.MustCompile content err")
		return
	}

	//content
	tmpContent := re2.FindAllStringSubmatch(result, -1) // -1 ,获取所有匹配的字符串
	for _, data := range tmpContent {
		content = data[1]
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "\t", "", -1)
	}

	return

}

// 把内容写入到文件
func StoreJoyToFile(i int, fileTitle, filteContent []string) {
	f, err := os.Create("joy/" + strconv.Itoa(i) + ".txt")
	if err != nil {
		fmt.Println("os.Create err: ", err)
		return
	}

	defer f.Close()

	n := len(fileTitle)
	for i := 0; i < n; i++ {
		f.WriteString(fileTitle[i] + "\n")
		f.WriteString(filteContent[i] + "\n")
		f.WriteString("*********************************************\n")
	}

}

func SpiderPage(i int, flag chan int) {
	url := "https://www.pengfue.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d页的内容: %s\n", i, url)

	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err: ", err)
		return
	}

	// fmt.Println("result = ", result)

	// 通过正则获取a标签的链接
	//<h1 class="dp-b"><a href="https://www.pengfue.com/content_1857807_1.html" target="_blank">会放坏</a>
	// 写正则的匹配规则,使用结构符 ``
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`) // ?s: 处理换行
	if re == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}

	// 获取到匹配的信息
	joyUrls := re.FindAllStringSubmatch(result, -1)
	// fmt.Println("joyUrls: ", joyUrls)

	// 创建切片保存title和content
	fileTitle := make([]string, 0)
	filteContent := make([]string, 0)

	// 获取a标签的url链接
	// 返回的第一个参数是下标，第二个是内容
	for _, data := range joyUrls {
		// fmt.Println("data = ", data[1])
		title, content, err := SpiderOneJoy(data[1])
		if err != nil {
			fmt.Println("SpiderOneJoy err")
			break
		}

		//fmt.Printf("title: #%v#\n", title)
		//fmt.Printf("content: #%v#\n\n", content)
		fileTitle = append(fileTitle, title)
		filteContent = append(filteContent, content)
	}
	// fmt.Println(fileTitle)
	// fmt.Println(filteContent)

	// 把内容写入到文件
	StoreJoyToFile(i, fileTitle, filteContent)
	flag <- i

}

func DoWork(start, end int) {
	flag := make(chan int)

	for i := start; i <= end; i++ {
		go SpiderPage(i, flag)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成！！！\n", <-flag)
	}
	// return
}

func main() {
	var start, end int
	fmt.Println("请输入起始页:( >=1)")
	fmt.Scan(&start)
	fmt.Println("请输入结束页:( >= 起始页)")
	fmt.Scan(&end)

	DoWork(start, end)

}
