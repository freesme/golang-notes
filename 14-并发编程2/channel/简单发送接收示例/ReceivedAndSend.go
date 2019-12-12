package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	//接收程序
	go func() {
		//在syncChan1中没接收到值之前会阻塞
		<-syncChan1
		fmt.Println("[Recevied-start] Recevied a sync signal and wait a second ...")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("[Recevied]", elem)
			} else {
				break
			}
		}
		fmt.Println("[Recevied=stop] Stooped")
		syncChan2 <- struct{}{}
	}()
	//发送程序
	go func() {
		defer close(strChan)
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("[Send] sent:", elem)
			//strChan中缓存了三个值，开启接收程序
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("[Send] Sent a sync signal")
			}
			//strChan最多缓存三个值  从strChan读取出"a"后, 发送程序再继续发送"d"
		}
		fmt.Println("[Send] wait 2 seconds.. ")
		time.Sleep(time.Second * 2)
		syncChan2 <- struct{}{}
	}()

	//发送和接收程序能完全执行syncChan2中会有两个值，在两次成功接收完成之前，会阻塞在h这里
	<-syncChan2
	<-syncChan2
}
