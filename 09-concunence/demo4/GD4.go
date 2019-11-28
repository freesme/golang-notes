package main

import "fmt"

func demo() {

	//声明channel var chanName chan 'ElementType'  ElementType指订这个channel所能传递的元素类型
	//eg: 声明一个传递类型为int 的channel
	var ch chan int
	//声明一个map，元素是bool型的channel
	var m map[string]chan bool
	//定义一个channel  声明并初始化一个 int 型的名为 ch 的channel
	ch := make(chan int)

	//写入
	ch <- value //向channel中写入数据通常会导致程序阻塞，直到有其他goroutine从channel中读取数据
	//读取
	value := <-ch //如果channel之前没有写入数据，那么从channel中读取数据也会导致程序阻塞，直到channel中被写入数据为止

	//之后还会提到如何控制channel只接受写或者只允许读取，即单向channel、

	// 3.带缓冲的channel  创建大小为1024的int类型的channel,没有读取方，写入方也可以一直往channel中写入，在缓冲区被填满之前都不会被阻塞
	c := make(chan int, 1024)

	//读取带缓冲的channel
	for i := range c {
		fmt.Println("Recived:", i)
	}

}
