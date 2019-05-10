package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello gooooo")
}

func main() {
	http.HandleFunc("/go", myHandler)

	// 在指定的地址进行监听，开启一个HTTP
	http.ListenAndServe("127.0.0.1:8000", nil)

	// 浏览器输入  127.0.0.1:8000/go
}
