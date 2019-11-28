package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	加锁的方式完成变量共享
*/
//10个goroutine中共享  变量counter
var counter int = 0

func Count(lock *sync.Mutex) {
	//加锁
	lock.Lock()
	counter++
	fmt.Println(counter)
	//释放锁
	lock.Unlock()

}
func main() {
	//定义一个锁
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock)
	}
	for { //检测 循环来不断检查counter的值 count=10 说明协程执行完成
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}

}
