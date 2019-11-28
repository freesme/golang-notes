package main

import "fmt"

//select的使用

/*
早在Unix时代， select机制就已经被引入。通过调用select()函数来监控一系列的文件句柄，一旦其中一个文件句柄发生了IO动作，
该select()调用就会被返回。

select的用法与switch语言非常类似，由select开始一个新的选择块，每个选择条件由
case语句来描述。与switch语句可以选择任何可使用相等比较的条件相比， select有比较多的
限制，其中最大的一条限制就是每个case语句里必须是一个IO操作
*/
func main() {
	runSelect1()
	//runSelect2()
}

func runSelect2() {
	ch := make(chan int, 1)
	for {
		select {
		case <-ch:
			fmt.Println("in there...")
		case ch <- 1:
		}
		i := <-ch
		fmt.Println("value received:", i)
	}
}

func runSelect1() {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <-ch
		fmt.Println("value received:", i)
	}
}
