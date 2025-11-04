package advanced

import (
	"fmt"
	"time"
)

//题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
//考察点 ：通道的基本使用、协程间通信。

func Job9() {
	ch := make(chan int, 5)
	go sendData(ch)
	go receiveData(ch)
	time.Sleep(time.Second * 5)
}

// 发送数据
func sendData(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		fmt.Println("发送数据：", i)
		ch <- i
	}
}

// 接收数据
func receiveData(ch <-chan int) {
	for {
		data := <-ch
		fmt.Println("接收数据：", data)
	}
}
