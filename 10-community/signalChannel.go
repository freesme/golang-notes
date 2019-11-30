package main

import (
	"fmt"
	"time"
)

/*
	单向通道演示

	单向通道通畅由双向通道转换而来，但是单向通道不可以转换回双向通道
	通道允许数据传递的方向是类型的一部分。
	对于两个通达类型而言，数据传递方向的不同就意味着他们的类型不同
*/

func main() {
	var strChan = make(chan string, 3)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go revice(strChan, syncChan1, syncChan2)

	go send(strChan, syncChan1, syncChan2)

	time.Sleep(time.Second * 10)
	<-syncChan2
	<-syncChan2

}

/*
	接收程序
*/
func revice(strChan <-chan string, syncChan1 <-chan struct{}, syncChan2 chan<- struct{}) {
	//随便接收个什么东西，不保存接收的内容 接收到内容之后会执行打印工作
	<-syncChan1
	fmt.Println("[接收] 接收一个单向sysn 等待1S")

	for {
		if elem, ok := <-strChan; ok {
			fmt.Println("[接收]", elem)
		} else {
			break
		}
	}
	fmt.Println("[接收]-停止")
	//向chan2随便发送什么
	syncChan2 <- struct{}{}
}

func send(strChan chan<- string, syncChan1 chan<- struct{}, syncChan2 chan<- struct{}) {
	for _, elem := range []string{"a", "b", "c", "d"} {
		//将数组内容发送到 strChan 中
		if elem == "c" {
			strChan <- elem
			fmt.Println("[发送]-[strChan]:", elem)
			time.Sleep(time.Second * 2)
		}
	}

	//关闭发送
	//close(strChan)
	//向chan2发送点什么
	syncChan1 <- struct{}{}
	fmt.Println("[发送]随便发点啥")
}
