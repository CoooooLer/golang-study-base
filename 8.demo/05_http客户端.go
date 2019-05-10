package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8000/index.html")
	if err != nil {
		fmt.Println("http.Get err: ", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("statucs = ", resp.Status)
	fmt.Println("StatusCoke = ", resp.StatusCode)
	fmt.Println("Header = ", resp.Header)
	// fmt.Println("Body = ", resp.Body)

	buf := make([]byte, 1024*4)
	var tmp string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body.Read err: ", err)
			break //切记是 break , 不是 return
		}
		tmp += string(buf[:n])
	}

	fmt.Println("tmp = ", tmp)
}
