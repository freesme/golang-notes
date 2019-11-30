package main

import (
	"fmt"
)

func main() {
	//定义一个chan int类型的双向通道
	var intChan chan int
	//定义只发送值的单向通道
	//var sendIntChan chan<- int
	//定义只接收值的单向通道
	var receiveChan <-chan int
	/*
		通道是在多个goroutine之间传递数据和同步的重要手段，而对通道操作的本身也是**同步**的 。在同一时刻，仅有一个goroutine能向一个通道发送元素值，
		同时也仅有一个goroutine能从它那里接收元素值。在通道中，各个元素值都是严格按照发送到此的先后顺序排列的，最早被发现送至通道的元素值会被先接收到。
		因此通道相当于一个FIFO(先进先出)的消息队列。此外通道通的元素值都具有原子性是不可被分割的，通道中的每一个元素值都可能被某一个goroutine接收，
		已被接收的元素值会立刻从通道中删除
	*/
	//初始化intChan 不带缓冲区的
	intChan = make(chan int)
	//有长度为1024的缓冲区
	//sendIntChan = make(chan int, 1024)
	receiveChan = make(chan int, 10)

	//接收元素值
	//从通达中接收值
	i := <-intChan
	fmt.Println(i)

	j, ok := <-receiveChan

	if ok {
		fmt.Println(j)
	} else {

	}

}
