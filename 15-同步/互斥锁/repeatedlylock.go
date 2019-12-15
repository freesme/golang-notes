package 互斥锁

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	fmt.Println("Lock the lock.(main)")

	mutex.Lock()
	fmt.Println("The lock is locked.(main)")
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println("Lock the lock ", i)
			mutex.Lock()
			fmt.Println("The lock is locked.", i)
		}(i)
	}
	time.Sleep(time.Second * 3)
	fmt.Println("Unlock the lock.(main)")
	mutex.Unlock()
	fmt.Println("The lock is unlocked.(main)")
	time.Sleep(time.Second * 3)
}

/*
Lock the lock.(main)
The lock is locked.(main)
Lock the lock  0
Lock the lock  1
Lock the lock  2
Unlock the lock.(main)
The lock is unlocked.(main)
//如果锁定了一个已锁定的互斥锁，那么进行重复锁定操作的goroutine将被阻塞，直到该互斥锁回到解锁状态
The lock is locked. 0
*/
