package demo

import "fmt"

func main() {
	//调用了10次Add()，应该有10次屏幕输出才对,但是此时屏幕上没有输出
	/*
			Go程序从初始化main package并执行main()函数开始，
			当main()函数返回时，程序退出，
			且程序并不等待其他goroutine（非主goroutine）结束

		要让主函数等待所有goroutine退出后再返回，如何知道goroutine都退出了呢?这就引出了多个goroutine之间通信的问题
	*/
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
}
func Add(x, y int) {
	z := x + y
	fmt.Println(x, "   ", z)
}
