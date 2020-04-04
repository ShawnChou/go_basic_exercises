package struct_test

import (
	"fmt"
	"testing"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

//会有多余的copy数据 通过unsafe.Pointer(e.Age)去查看的时候两次地址不一样
func (e Employee) String() string {
	return fmt.Sprintf("Id is %s,Name is %s,age is %d", e.Id, e.Name, e.Age)
}

func (e *Employee) String() string {
	return fmt.Sprintf("Id is %s,Name is %s,age is %d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"1", "Shawn", 2}
	fmt.Print(e.String())
}
