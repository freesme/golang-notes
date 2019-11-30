package main

import (
	"fmt"
	"time"
)

/*
	关闭通道相关

	无论怎样都不应该在接收端关闭通道，因此在接收端通常无法判断接收端是否还会想该通道发送元素值
	另一方面，在发送端关闭通道一般不会对接收端的接收操作产生什么影响，如果通道在被关闭时其中仍有元素值，
	你依然可以用接收表达式取出，并根据表单时的第二个结果值（elem,ok := <- typeChan）判断通道是否已关闭且无元素值可以取出



*/

func main() {
	dataChan := make(chan int, 5)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() {
		<-syncChan1
		for {
			if elem, ok := <-dataChan; ok {
				fmt.Println("接收:", elem)
			} else {
				break
			}
			fmt.Println("[接收完成]")
			syncChan2 <- struct{}{}
		}

	}()

	go func() {
		for i := 0; i < 5; i++ {
			dataChan <- i
			fmt.Println("[发送].", i)
		}

		close(dataChan)
		syncChan1 <- struct{}{}
		fmt.Println("[发送完成]")
		time.Sleep(5 * time.Second)
		syncChan2 <- struct{}{}
	}()

	<-syncChan2
	<-syncChan2
	time.Sleep(10 * time.Second)
}
