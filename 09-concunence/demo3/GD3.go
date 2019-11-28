package main

import "fmt"

/*
		channel是Go语言在语言级别提供的goroutine间的通信方式。我们可以使用channel在两个或
	多个goroutine之间传递消息。 channel是进程内的通信方式，因此通过channel传递对象的过程和调
	用函数时的参数传递行为比较一致，比如也可以传递指针等。

	(如果需要跨进程通信，我们建议用分布式系统的方法来解决，比如使用Socket或者HTTP等通信协议。
	Go语言对于网络方面也有非常完善的支持)

	channel是类型相关的，一个channel只能传递一种类型的值，这个类型需要在声明channel的时候指定
	(如果对Unix管道有所了解的话，就不难理解channel，可以将其认为是一种类型安全的管道)
*/

//改造 GD2.go
func Count(ch chan int) {
	//向对应的channel中写入一个数据,在对应的channel被读取之前这个操作时阻塞的
	ch <- 1
	fmt.Println("Counting")
}

func main() {
	//定义10个长度的channel的数组
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		//给数组的第 i 位赋值
		chs[i] = make(chan int)
		//把10个channel分配给 10个goroutine
		go Count(chs[i])
	}

	for _, ch := range chs {
		//所有goroutine启动完成后，通过 <-ch语句从10个channel中依次读取数据，在对应的channel写入数据前
		//这个操作也是阻塞的，这样就用channel实现了类似锁的功能，从而保证了所有goroutine完成后主函数才返回
		<-ch
		fmt.Println("<-ch")
	}
}
