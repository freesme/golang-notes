package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		字面常量 -> 硬编码的常量
		-12			整数类型
		3.14159265358979323846	浮点类型
		3.2+12i		复数类型
		true		布尔类型
		"zhangsan" 字符串类型

	*/
	// 1.常量定义 const关键字
	const Pi float64 = 3.14159265358979323846
	const zero = 0.0 //无类型浮点常量
	const (
		size int64 = 1024
		eof        = -1 //无类型整形变量
	)
	const u, v float64 = 0, 3 //多重赋值 u = 0.0  v=3.0

	fmt.Println(reflect.TypeOf(u), "---", reflect.TypeOf(v)) //查看常量类型
	fmt.Println(u, "----", v)

	//常量可以是一个在编译期运算的常量表达式，不能出现运行期才能得到结果的biaodashi
	const mask = 1 << 3 //编译期

	//const Home = os.Getenv("HOME")	//报错

	//	2.预定义常量
	/*
		Go预定义了 true false和iota
		iota可以被认为是一个可被编译器修改的常量，在每一个const关键字出现的时候被重置为 0
		然后在下一个const初心之前，每出现一次iota 其所代表的的数字就会自动增加 1
	*/
	const (
		c0 = iota //c0 == 0
		c1 = iota //c1 == 1
		c2 = iota //c2 == 2
	)

	const (
		a = 1 << iota //a == 1
		b = 1 << iota //b == 2
		c = 1 << iota //c == 4
	)

	const (
		z = iota * 42 // z = 0
		x = iota * 42 // x = 42.0
		w = iota * 42 // w = 84.0
	)

	//	3.枚举  -Go语言不支持众多其他语言明确支持的enum 关键字
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)

	fmt.Println(Sunday)
	fmt.Println(numberOfDays)

}
