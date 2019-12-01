package main

import "fmt"

/*
	select 与 for 语句连接
*/

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//
	close(intChan)
	syncChan := make(chan struct{}, 1)
	go func() {
	Loop:
		for {
			select {
			//读取出通道中所有值
			case e, ok := <-intChan:
				if !ok {
					fmt.Println("END.")
					//结束for循环
					break Loop
				}
				fmt.Printf("Received:%v\n", e)
			}
		}
		syncChan <- struct{}{}
	}()
	<-syncChan
}
