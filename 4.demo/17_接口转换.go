package main

import "fmt"

type Humaner interface {
	Say()
}

type Personer interface {
	Humaner
	Sing(lrc string)
}

type Student struct {
	id   int
	name string
}

func (tmp *Student) Say() {
	fmt.Printf("Student[%d, %s] say \n", tmp.id, tmp.name)
}

func (tmp *Student) Sing(lrc string) {
	fmt.Printf("Student[%d, %s] sing %s\n", tmp.id, tmp.name, lrc)
}

func main() {
	// 超集可以转换为子集，但是子集不能转换为超集
	var i Humaner     // 子集
	var iPro Personer // 超集

	s := Student{12345, "jannnnne"}
	i = &s
	iPro = &s

	i.Say()
	iPro.Say()

	// i.Sing("iiii")  // error
	// iPro.Sing("pro")

	i = iPro
	// iPro = i // error

}
