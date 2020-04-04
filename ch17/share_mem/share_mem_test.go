package share_mem

import (
	"sync"
	"testing"
	"time"
)

//结果不到5000 导致并发竞争条线，非线程安全的程序。
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
	//t.Log(counter)
}

//对共享的内存进行锁的保护  线程安全
func TestCounter2(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
	//t.Log(counter)
}

//这里的时间是约定的，不一定准确，让协程结束可通知
func T2stCounter3(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait() //等待所有协程结束
	t.Logf("counter = %d", counter)
	//t.Log(counter)
}
