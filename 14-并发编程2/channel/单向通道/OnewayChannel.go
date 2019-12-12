package main

import (
	"fmt"
	"time"
)

/**
无论哪一种单向通道，都不应该出现在变量的声明中  单向通道应该有双向通通道变换而来

	func Notify(c chan<- os.Signal.sig ...os.Signal)
该函数的第一个参数的类型是发送通道类型，从表面上看，调用它的程序应该需要传入一个只能发送而不能接收通道，实际上应该出入一个双向
通道，Go会根据该参数的声明，自动把它转为单向通道。		该函数中的代码只能向通道c发送值，而不能从它那里接收值，这是一个
强约束，从该函数中的c中获取值会造成编译错误，在函数之外c不受约束
既然Notify函数中的代码只能对它进行发送操作，那么函数外的代码只应对它进行接收操作。函数外的发送操作只会造成干扰

*/

var strChan = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go receive(strChan, syncChan1, syncChan2)
	go send(strChan, syncChan1, syncChan2)

	<-syncChan2
	<-syncChan2
}

/*
	strChan 接收类型的通道，函数内部只能从中获取值
*/
func receive(strChan <-chan string, syncChan1 <-chan struct{}, syncChan2 chan<- struct{}) {
	//等待发送方发送一个开始的信号
	<-syncChan1
	fmt.Println("[RECEIVER] Received a sync signal and wait second...")
	time.Sleep(time.Second)
	for {
		if elem, ok := <-strChan; ok {
			fmt.Println("[RECEIVER]:", elem)
		} else {
			break
		}
	}
	fmt.Println("[RECEIVER]:Stopped")
	syncChan2 <- struct{}{}
}

/*
	发送函数 只能对strChan发送操作
*/
func send(strChan chan<- string, syncChan1 chan<- struct{}, syncChan2 chan<- struct{}) {

	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		fmt.Println("[SEND]", elem)
		if elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("[SEND]Sent a sync signal")
		}
	}
	defer close(strChan)

	fmt.Println("[SEND] wait 2 seconds")
	time.Sleep(time.Second * 2)

	syncChan2 <- struct{}{}
}
