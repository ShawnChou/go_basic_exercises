package interface_test

import "testing"

//与其他编程语言的差异
//1.接口为非入侵性的，实现不依赖于接口
//2.所以接口的定义可以包含在接口使用者的包内
type Programer interface {
	WriteHelloWorld() string
}

type GoProgramer struct {
}

func (p *GoProgramer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestProgramer(t *testing.T) {
	var p Programer
	p = new(GoProgramer)
	t.Log(p.WriteHelloWorld())
}

//接口变量
var prog Programer = &GoProgramer{}
