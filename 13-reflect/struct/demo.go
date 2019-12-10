package main

import (
	"fmt"
	"reflect"
)

/*
通过反射获取结构体对象的属性和方法
通过反射调用结构体的方法
*/

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("Say: ", msg)
}
func (p Person) PrintInfo() {
	fmt.Println("姓名：", p.Name, "年龄：", p.Age, "性别：", p.Sex)
}
func main() {
	p1 := Person{"zhang", 32, "man"}
	//获得对象类型
	t1 := reflect.TypeOf(p1)
	fmt.Println("Type of p1:", t1)
	fmt.Println("p1的类型：", t1.Name())

	//返回该类型的特定类型 struct
	k1 := t1.Kind()
	fmt.Println(k1)

	v1 := reflect.ValueOf(p1)
	fmt.Println(v1)

	if t1.Kind() == reflect.Struct {
		//是结构体类型，获取里面的字段名字
		fmt.Println(t1.NumField()) //3
		for i := 0; i < t1.NumField(); i++ {
			field := t1.Field(i)
			fmt.Println(field)             //{Name  string  0 [0] false},{Age  int  16 [1] false},{Sex  string  24 [2] false}
			val := v1.Field(i).Interface() //通过interface方法来取出这个字段所对应的值
			fmt.Printf("字段名字：%s,字段类型：%s,字段数值：%v\n", field.Name, field.Type, val)
		}
	}
	//2.操作方法
	for i := 0; i < t1.NumMethod(); i++ {
		m := t1.Method(i)
		fmt.Println(m.Name, m.Type) //Hello func(main.Person)
		/*
		   {Hello  func(main.Person) <func(main.Person) Value> 0}
		   {PrintInfo  func(main.Person) <func(main.Person) Value> 1}
		*/
	}
	m1 := v1.MethodByName("Say")
	args := []reflect.Value{reflect.ValueOf("干啥呢？")}
	m1.Call(args)

	m2 := v1.MethodByName("PrintInfo")
	m2.Call(nil)

}
