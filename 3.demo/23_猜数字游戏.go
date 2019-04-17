package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var randNum int
	RandNum(&randNum) //获取 randNum 的地址，以指针方式赋值
	// fmt.Println("randNum = ", randNum)

	RandSlice := make([]int, 4) // 使用 make 函数声明一个长度、容量都为4的切片
	GetNum(RandSlice, randNum)
	fmt.Println("RandSlice = ", RandSlice)

	OnGame(RandSlice)
}

//产生随机数
func RandNum(p *int) { // 指针方式
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num = rand.Intn(10000) // 取 10000 以内的数，不包含 10000
		if num > 1000 {
			break
		}
	}
	*p = num // *p 指向 randNum
}

// 取出 num 的每一位数放入 s 切片中
func GetNum(s []int, num int) {
	s[0] = num / 1000       // 千位
	s[1] = num % 1000 / 100 // 百位
	s[2] = num % 100 / 10   // 十位
	s[3] = num % 10         // 个位
}

func OnGame(RandSlice []int) {
	var num int                // 声明变量，保存用户输入的数
	keySlice := make([]int, 4) //使用 make 函数声明一个切片
	fmt.Println("请输入一个4位整数：")

	for {
		for {
			fmt.Scan(&num)
			GetNum(keySlice, num)
			// fmt.Println("keySlice = ", keySlice)
			if 999 < num && num < 10000 {
				break
			}
			fmt.Println("输入的数不符合规则,请重新输入：")
		}
		f := 0
		for i := 0; i < 4; i++ {
			if keySlice[i] > RandSlice[i] {
				fmt.Printf("第%d位大了一点点,", i+1)
			} else if keySlice[i] < RandSlice[i] {
				fmt.Printf("第%d位小了一点点,", i+1)
			} else {
				f++
				fmt.Printf("第%d位猜对了。", i+1)
			}
		}
		if f == 4 {
			fmt.Println("全部猜对了！！！")
			break
		} else {
			fmt.Println("请再猜：")
		}
	}

}
