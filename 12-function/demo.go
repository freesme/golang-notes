package main

import (
	"fmt"
	"math"
)

//定义一个函数
func funcName(parametername1 string, parametername2 int) (output1 int, output2 string) {
	return parametername2, parametername1
}

//可变参数
func addfunc(arg ...int) int {
	v := 0
	for _, i := range arg {
		v += i
	}
	return v
}

//引用传递
func add(a *int) {
	*a = *a + 1
}

//回调函数使用
func printRes(fun int) {
	fmt.Println("回调函数调用结果:", fun)
}

//闭包
/*
一个外层函数中，有内层函数，该内层函数中，会操作外层函数中的局部变量(外层函数中的参数，
或者外层函数直接定义的变量)，并且该外层函数的返回值就是这个内层函数。
该内层函数和外层函数的局部变量，统称为闭包结构。
局部变量的生命周期会发生改变，正常的局部变量随着函数的调用而创建，随着函数的结束而销毁。
但是闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁，因为内层函数还要继续使用。
*/
func increment() func() int { //外层函数
	//1.定义了一个局部变量
	i := 0

	//2.定义了一个匿名函数，给变量自赠了并返回
	fun := func() int { //内层函数
		i++
		return i
	}
	//3.返回该匿名函数
	return fun
}

func main() {
	//延迟函数
	defer fmt.Println("延迟函数1 最后运行")
	defer fmt.Println("延迟函数运行顺序")

	//使用函数
	age, name := funcName("lisi", 12)
	fmt.Println(name, age)

	//值传递
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	//引用传递
	var x int = 8
	add(&x)
	fmt.Println(x)

	//匿名函数
	res := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println("匿名函数运行结果:", res)
	//回调函数调用
	printRes(addfunc(1, 2, 4))

	//闭包
	f := increment()
	o := f()
	fmt.Println("闭包：", o)
}
