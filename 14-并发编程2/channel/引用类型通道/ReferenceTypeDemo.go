package main

import (
	"fmt"
	"time"
)

/*
发送方向通道发送的值会被复制，接收方接收的总是该值的副本，而不是该值本身
当接收方从通道接收到一个值类型的值时，对该值得修改就不会影响到发送方持有的那个源值。但对于引用类型的值来说，这种修改会同时影响收发
方持有的值
*/

var mapChan = make(chan map[string]int, 1)

func main() {
	syncChan := make(chan struct{}, 2)
	//演示接收操作
	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				elem["count"]++
			} else {
				break
			}
		}
		fmt.Println("[RECEIVER] Stopped")
		syncChan <- struct{}{}
	}()

	//演示发送操作
	go func() {
		defer close(mapChan)
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			//将countMap发送了5次，接收程序中对这个map "count" ++
			mapChan <- countMap
			//休眠过程 接收方法已经修改过了 map[count]
			//mapChan的元素类型属于引用类型，因此，接收对象元素值得副本的修改会影响到发送方持有的源值
			time.Sleep(time.Millisecond)
			fmt.Println("[SEND]The count map", countMap)
		}
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}
