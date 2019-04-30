package main

import (
	"fmt"
	"regexp"
)

func main() {
	// `` 原生字符串
	buf := `<body>
	<div class="bigbigbox">
         <div clasee="box"></div>
	</div>
	<ul>
		<li>1111</li>
		<li>1111</li>
		<li>1111</li>
		<li>1111</li>
		<li>1111</li>
		<li>1111</li>
		<li>1111</li>
		<li>1111</li>
		<li>1111</li>
		<li>1111</li>
		
	</ul>
	<div class="box">
		<div>1</div>
		<div>2
		2
		2
		3
		</div>
		<div>1</div>
		<div>3</div>
		<div>4</div>
	</div>
</body>`

	// reg := regexp.MustCompile(`<div>(.*)</div>`)
	reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`) // 匹配换行

	res := reg.FindAllStringSubmatch(buf, -1)
	// fmt.Println(res)

	// 过滤 <div></div>
	for _, data := range res {
		// fmt.Println(data[0]) // 带<div></div>
		fmt.Println(data[1]) // 不带 <div></div>
	}
}
