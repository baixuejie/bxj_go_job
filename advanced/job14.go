package advanced

import (
	"fmt"
	"sync"
)

func Job14() {
	//创建一个缓冲区为3的通道
	ch := make(chan int, 3)
	ws := sync.WaitGroup{}
	ws.Add(1)
	go func() {
		defer ws.Done()
		for i := 0; i < 1000; i++ {
			//发送数据
			ch <- i
			fmt.Println("send data:", i)
		}
		close(ch)
	}()
	ws.Add(1)
	go func() {
		defer ws.Done()
		for num := range ch {
			fmt.Println("receive data:", num)
		}
	}()
	ws.Wait()
}
