package main

import (
	"fmt"
	"time"
)

func main() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	//forgo1(names)
	forgo2(names)

	time.Sleep(time.Millisecond) //让go程在 main(主goroutine)执行完前有时间执行完成

}

/*
	输出结果与预想的有差距
	Hello  Mark
	Hello  Mark
	Hello  Mark
	Hello  Mark
	Hello  Mark

	在这里并发执行的5个go函数中，name的值都是Mark。这是因为他们都是在for语句执行完毕后才执行的，
而name在那时指代的值已经是Mark了
*/
func forgo1(names []string) {
	for _, name := range names {
		go func() {
			fmt.Println("Hello ", name)
		}()
	}
}

/*
	执行出 预想中正确的结果
	Hello  Eric
	Hello  Harry
	Hello  Robert
	Hello  Jim
	Hello  Mark
*/
func forgo2(names []string) {
	for _, name := range names {
		go func() {
			fmt.Println("Hello ", name)
		}()

		time.Sleep(time.Millisecond) //每次迭代完成 让go程有时间执行
	}
}
