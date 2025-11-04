package advanced

import (
	"sync"
	"sync/atomic"
)

func Job16() {

	ms := sync.WaitGroup{}
	count := int32(0)
	for i := 0; i < 10; i++ {
		ms.Add(1)
		go func() {
			defer ms.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&count, 1)
			}
		}()
	}
	ms.Wait()
	println("最终数量为：", count)
}
