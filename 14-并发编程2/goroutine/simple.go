package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//goDemo()
	//goDemo2()

	//goDemo3()
	//goDemo4()
	goDemo5()

}

func goDemo() {
	go println("go 程执行")
	//time.Sleep(time.Millisecond)  //main线程会在go程执行前完成，看不到go程执行的信息
	runtime.Gosched() //暂停当前的Goroutine 让其他的G有机会运行  在实际的复杂情况下此方法也不适用
}

func goDemo2() {
	name := "Eric"
	go func() {
		fmt.Println(name)
	}()
	//time.Sleep(time.Millisecond)
	name = "Harry"
	time.Sleep(time.Millisecond)
}

//令人迷惑的结果
func goDemo3() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func() {
			fmt.Println(name)
		}()
	}
	time.Sleep(time.Millisecond)
}

func goDemo4() {
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func() {
			fmt.Println(name)
		}()
		//每次迭代完成之前基于之前go函数一个执行的机会
		time.Sleep(time.Millisecond)
	}
}

//可冲入的方法   让go函数中使用的name的值不会受到外部变量变化的影响
func goDemo5() {
	//把迭代变量的值作为参数传递给go函数
	names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
	for _, name := range names {
		go func(name string) {
			fmt.Println(name)
		}(name)
	}
	time.Sleep(time.Millisecond)
}
