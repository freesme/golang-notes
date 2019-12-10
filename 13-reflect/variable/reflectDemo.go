package variable

import (
	"fmt"
	"reflect"
)

/*
反射
反射机制是程序能够检查自身结构，属于元变成的范畴
Java的反射机制
我们所熟知的Java的反射机制是什么？对于类和对象的使用，普通的方式是知道类和对象的属性和方法之后进行调用或者访问。
而反射机制，简单来说，是在运行状态中，Java对于任何的类，都能够确认到这个类的所有方法和属性；对于任何一个对象，都能调用它的任意方法和属性。这种动态获取或者调用的方式就是Java的反射机制。

能做什么
在Java中，通过反射机制在运行时能够做到如下：

确认对象的类
确认类的所有成员变量和方法
动态调用任意一个对象的方法
*/

func main() {
	//1.接口类型变量 => 反射类型对象
	var circle float64 = 6.28
	var icir interface{}

	icir = circle
	fmt.Println("Reflect : circle.Value =", reflect.ValueOf(icir)) //Reflect : circle.Value =  6.28
	fmt.Println("Reflect : circle.Type  = ", reflect.TypeOf(icir)) //Reflect : circle.Type =  float64

	// 2. 反射类型对象 => 接口类型变量
	v1 := reflect.ValueOf(icir)
	fmt.Println(v1)
	fmt.Println(v1.Interface())

	y := v1.Interface().(float64)
	fmt.Println(y)

	//3.修改
	fmt.Println(v1.CanSet())       //是否可以进行修改
	v2 := reflect.ValueOf(&circle) //传递指针才能修改
	v4 := v2.Elem()
	fmt.Println(v4.CanSet())

	v4.SetFloat(3.14)
	fmt.Println(circle)

}
