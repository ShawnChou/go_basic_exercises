package obj_cache_test

// 适合于通过复用，降低复杂对象的创建和GC代价
// 协程安全，会有锁的开销
// 生命周期收GC影响，不适合于做连接池等，需自己管理生命周期的资源的池化

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

//与channel方式的区别
//GC会清除sync.pool缓存对象
//缓存对象存在于下次GC之前

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	runtime.GC() //手动触发GC 结果重新调用了上面的New
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
	v2, _ := pool.Get().(int)
	fmt.Println(v2)
}

func TestSuncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object")
			return 10
		},
	}
	pool.Put(100)
	pool.Put(100)
	pool.Put(100)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			t.Log(pool.Get())
		}(i)
	}
	wg.Wait()
}
