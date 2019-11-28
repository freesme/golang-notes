package main

import "fmt"

func main() {
	//	1.变量声明
	//----------方式一----------
	var v1 int
	var v2 string
	var v3 [10]int //数组
	var v4 []int   //数组切片
	var v5 struct {
		f int
	}
	var v6 *int
	var v7 map[string]int //键为string类型 值为int类型

	//----------方式二----------
	var (
		v11 int
		v12 string
	)

	//	2.变量初始化
	var v21 int = 10
	var v22 = 10 //编译器可以自动推导v22的类型
	v23 := 10    //编译器可以自动推导v23的类型

	// := 左边的变量不应该是已经被声明过的
	var i int
	i := 2 //编译错误 no new variables on left side of :=

	//	3.变量赋值
	var v30 int
	v30 = 10

	j := 2
	//多重赋值功能 eg 交换变量
	i, j = j, i

	//	3.匿名变量
	_, lastName, _ := getName()
	fmt.Println(lastName)
}

func getName() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibo Marruko"
}
