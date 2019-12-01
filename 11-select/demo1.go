package main

import "fmt"

/*
	分支选择规则  所有跟在case关键字右边的发送语句或接收语句中的通道表达式和原色表达式都会先求值（求值的顺序是从左到右从上而下的）无论他们所在的case
	是否有可能被选择都会是这样

	一个select只能有一个default分支，不过它的位置可以放在改语句的任何位置上
*/

var intChan1 chan int
var intChan2 chan int
var channels = []chan int{intChan1, intChan2}
var numbers = []int{1, 2, 3, 4, 5}

func main() {
	select {
	default:
		fmt.Println("Default.")
	case getChan(0) <- getNumber(0):
		fmt.Println("1th case is selected")
	case getChan(1) <- getNumber(1):
		fmt.Println("the 2nd case is selected")

	}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}
