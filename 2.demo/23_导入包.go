package main

// 方式1：导入包
// import "fmt"  // 导入包必须使用，否则会报错
// import "os"

// 方式2：导入包
// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Println("111111")
// 	fmt.Println("os.Args = ", os.Args)
// }

// . 操作
// import . "fmt" // 调用函数，无需通过包名
// import . "os"

// func main() {
// 	Println("22222222")
// 	Println("简写的os.Args = ", Args)
// }

// 给包取别名
// import io "fmt"

// func main() {
// 	io.Println("给包起一个别名")
// }

// 忽略此包
import _ "fmt"

func main() {

}
