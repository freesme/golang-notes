package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	//用于演示接收操作
	go func() {
		//试图从syncChan1接收信号，在发送方发送完这个信号之前会阻塞
		<-syncChan1

		fmt.Println("Received a sync and wait a second... [receiver]")
		time.Sleep(time.Second)

		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("Receiver:", elem, "[received]")
			} else {
				break
			}
		}
		fmt.Println("Stopped.[receiver]")

		syncChan2 <- struct{}{}
	}()

	//演示发送操作
	go func() {
		for _, elem := range []string{"a", "b", "c", "d", "e"} {
			strChan <- elem
			fmt.Println("Sent:", elem, "[sender]")
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("sent a sync signal.[sender]")
			}
			fmt.Println("wait 2 seconds... [sender]")
			time.Sleep(time.Second * 2)
			//关闭通道  不关闭会死锁，关闭程序执行效果不理想
			//close(strChan)
		}
		syncChan2 <- struct{}{}
	}()

	<-syncChan2
	<-syncChan2
}

/*
	Sent: a [sender]
	wait 2 seconds... [sender]
	Sent: b [sender]
	wait 2 seconds... [sender]
	Sent: c [sender]
	sent a sync signal.[sender]
	wait 2 seconds... [sender]
	Received a sync and wait a second... [receiver]
	Receiver: a [received]
	Receiver: b [received]
	Receiver: c [received]
	Receiver: d [received]
	Sent: d [sender]
	wait 2 seconds... [sender]
	Sent: e [sender]
	wait 2 seconds... [sender]
	Receiver: e [received]

	fatal error: all goroutines are asleep - deadlock!
*/
