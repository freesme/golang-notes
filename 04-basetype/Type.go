package main

import (
	"fmt"
	"math"
)

func main() {
	//	1.bool
	var v1 bool
	v1 = true
	v2 := (1 == 2)
	fmt.Println(v1, "____", v2)
	// bool 不能接收其他类型的赋值 不支持自动或者强制的类型转换
	//var b bool
	//b = 1 	//错误
	//b = bool(1)
	//	正确用法
	var b bool
	b = (1 != 0)
	fmt.Println("Result: ", b)

	//	2.整型 			长度		    范围
	//	int 			  1			-128 ~ 127
	//	unit8(= byte) 	  1			0 ~ 255
	//	int16 			  2			-32,768 ~ 32767
	//	uint16 			  2			0 ~ 65535
	//	int32 			  4			-2,147483,648 ~ 2,147483,647
	//	uint32 			  4			0 ~ 4,294,967,295
	//	int64 			  8			-9,223,372,036,854,775,808 ~ 9,223,372,036,854,775,807
	//	uint64			  8			0 ~18,446,744,073,709,551,615
	//int uint 长度和值范围与运行平台有关  intptr 长度同指针 值范围 x86 4字节 x64 8字节

	//	2.1类型转换   int 和 int32是两种不同的类型编译器不会自动做类型转换

	var value2 int32
	value1 := 64
	value2 = int32(value1)
	fmt.Println(value2)

	// 	2.2 运算符
	fmt.Println("取余运算:", 5%3) // =2

	//比较运算  >,<,==,>=,<=,!=
	//位运算  << 左移,>> 右移,x^y 异或 ,& 与,| 或,^x 取反

	//	3.浮点型
	//	3.1浮点数比较不是一种精确的比较， 使用== 是不稳定的
	f1 := 0.1000123333
	f2 := 0.200112234
	p := 0.000001
	equal := IsEqual(f1, f2, p)

	fmt.Println(f1, " < ", f2, "?:", equal)

	// 4.字符串
	var str string
	str = "Hello World"
	ch := str[0]

	fmt.Printf("The length of\"%s\" is %d .\n", str, len(str)) //获取字符串长度
	fmt.Printf("The first char of\"%s\" is %c.\n", str, ch)

	//str[0] ='x'		//编译错误 字符串的内容不能在初始化后被修改

	//	5.数组
	//[32]byte					//长度为32的数组，每个元素为一个字节
	//[2*N] struct{x,y int32} 	//复杂类型数组
	//[1000]*float64			//指针数组
	//[3][5]int					//二维数组
	//[2][2][2]float64 			//等同于[2]([2]([2]float64))
	//	5.1元素访问
	var array = [10]int{1, 2, 3, 45, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(array); i++ {
		fmt.Println("array[", i, "]", array[i])
	}

	//range

	fmt.Println("---------使用RANGE方式遍历---------")
	for i, v := range array {
		fmt.Println("array[", i, "]", v)
	}

	//	5.2值类型	Go语言中数组是一个值类型，所有值类型变量在  赋值和作为参数传递  时将产生一次复制动作
	// 如果将数组作为函数的参数类型，则在函数调用时该参数将发生数据复制，因此在函数体重无法修改传入的数组的内容，
	// 函数内操作的知识所传入数组的一个副本

	modify(array)
	fmt.Println("实际:", array)

	// 5.3数组切片 slice  用来弥补 值类型的不足
	/*
		切片的数据结构
			1.指向原生数的指针
			2.数组切片中的元素个数
			3.数组切片已分配的存储空间
	*/
	fmt.Println("---------------数组切片 slice ---------------")
	// 创建切片
	// 1.基于数组  语法 array[first:last]  所有元素[:] 前5个 [:5] 从第5个到最后[5:]
	var mySlice1 []int = array[:5]
	rangeArr(mySlice1)
	//2.直接创建  make
	//mysile2 := make([]int, 5)     //创建一个初始元素个数为5的数组切片，元素初始值为0
	//mySile3 := make([]int, 5, 10) //直接创建并初始化包含5个元素，初始值 0 ，预留10个存储空间
	mySile4 := []int{1, 2, 3, 4, 5}
	mySile4 = append(mySile4, 1, 2, 3)

	rangeArr(mySile4)
	fmt.Println("实际长度len:", len(mySile4))
	fmt.Println("预留长度cap:", cap(mySile4))

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := []int{6, 7, 8}

	copy(slice1, slice2) //只会复制slice1的前三个到slice2
	copy(slice4, slice3) //只会复制slice4的前三个到slice3的前3个位置

	//	6.map
	fmt.Println("--------------- map ---------------")
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	//往这个map里插入几条数据
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 201"}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 202"}

	person, ok := personDB["1"]
	if ok {
		fmt.Println("Found Person", person.Name, "with id:", person.ID, "room:", person.Address)
	} else {
		fmt.Println("查无此人")
	}

	//删除元素
	delete(personDB, "1")
	person1, ok := personDB["1"]
	if ok {
		fmt.Println("Found Person", person1.Name, "with id:", person1.ID, "room:", person1.Address)
	} else {
		fmt.Println("查无此人")
	}
}

// p为用户自定义精度
func IsEqual(f1, f2, p float64) bool {
	return math.Dim(f1, f2) < p
}

//修改数组值
func modify(array [10]int) {
	array[0] = 10
	fmt.Println("副本:", array)
}

// 遍历切片
func rangeArr(arr []int) {
	fmt.Println("遍历切片", arr)
	for i, v := range arr {
		fmt.Println("index:", i, "value:", v)
	}
}

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}
