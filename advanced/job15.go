package advanced

import "sync"

func Job15() {

	ms := sync.WaitGroup{}
	mx := sync.Mutex{}
	count := 0
	for i := 0; i < 10; i++ {
		ms.Add(1)
		go func() {
			defer ms.Done()
			for j := 0; j < 1000; j++ {
				mx.Lock()
				count++
				mx.Unlock()
			}
		}()
	}
	ms.Wait()
	println("最终数量为：", count)

}
