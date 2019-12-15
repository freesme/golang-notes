package main

import (
	"fmt"
	"time"
)

/*
	非缓冲的channel
	如果在初始化一个通道时将其容量设置成0，或者直接忽略对容量的设置，就会使该通道成为一个非缓冲通道，与
	以异步的方式传递元素值的缓冲通道不同，非缓冲通道只能同步地传递元素值

	与缓冲通道相比，针对非缓冲通道的happens before的原则有两个特别之处
	1.向此类通道发送元素值的操作会被阻塞，直到至少有一个针对该通道的操作接收进行位置，该接收操作会先得到元素值的副本，然后在唤醒发送方所在的goroutine
之后返回，也就是说，这类操作会在对应的发送操作完成之前完成
	2.从此类通道接收元素值的操作会被阻塞，直到至少有一个针对该通道的操作发送进行位置，该接收操作会直接把元素值复制给对方，然后在唤醒接收方所在的goroutine
之后返回，也就是说，这类操作会在对应的接收操作完成之前完成


同步特性：由于非缓冲通道会以同步的方式传递元素值，在其上收发元素值的速度总是与慢的哪一方持平
*/

func main() {
	sendingInterval := time.Second
	receptionInterval := time.Second * 2

	intChan := make(chan int, 0)

	go func() {
		var ts0, ts1 int64
		for i := 1; i < 5; i++ {
			intChan <- i
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("Sent:", i)
			} else {
				fmt.Printf("Sent:%d [interval:%d s]\n", i, ts1-ts0)
			}
			ts0 = time.Now().Unix()
			time.Sleep(sendingInterval)
		}
		close(intChan)
	}()
	var ts0, ts1 int64

Loop:
	for {
		select {
		case v, ok := <-intChan:
			if !ok {
				break Loop
			}
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("Reveived:", v)
			} else {
				fmt.Printf("Received:%d[interval: %d s]\n", v, ts1-ts0)
			}
		}
		ts0 = time.Now().Unix()
		time.Sleep(receptionInterval)
	}
	fmt.Println("END.")
}

//发送操作和接收操作的时间间隔都与receptionInterval变量一直
