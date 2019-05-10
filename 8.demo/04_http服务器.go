package main

import (
	"fmt"
	"net/http"
)

// w 回复客户端
// r 获取客户端发送的数据
func HandConn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.Mehtod = ", r.Method)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Header = ", r.Header)
	fmt.Println("r.Body = ", r.Body)

	w.Write([]byte("hello goooo"))
}

func main() {
	// 注册处理函数，用户连接，自动调用指定的处理函数
	http.HandleFunc("/index.html", HandConn)

	// 监听绑定
	http.ListenAndServe(":8000", nil)
}
