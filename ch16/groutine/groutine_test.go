package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		//gp的参数传递是值传递
		go func(i int) {
			fmt.Println(i)
		}(i)
		// i 在携程里共享了，共享变量就要有竞争条件，一般要用锁来完成
		//go func() {
		//fmt.Println(i)
		//}()
	}
	time.Sleep(time.Millisecond * 50)
}
