package main

import (
	"fmt"
)

//	1.为类型添加方法
type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

//在需要修改对象的时候 需要使用指针
func (a *Integer) Add(b Integer) {
	*a += b
}

/*
结构体
	Go语言结构体 struct 和其他语言的 class具有同样的地位，但是Go语言放弃了包括继承在内的大量面向对象特性，
	只保留组合(Composition)这个基础的特性
*/

type Rect struct {
	x, y          float64
	width, height float64
}

//go语言中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以NewXXX来命名
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

//匿名组合
type Base struct {
	Name string
}

func (b *Base) Foo() {
	fmt.Println("Base::Foo()")
}
func (b *Base) Bar() {
	fmt.Println("Base::Bar()")
}

type Foo struct {
	Base
}

//继承自Base
func (foo *Foo) Bar() {
	foo.Base.Bar()
}

func (foo *Foo) Foo() {
	fmt.Println("Foo::Foo()")
}

func main() {
	var a Integer = 1
	if a.Less(3) { // 面向对象
		fmt.Println("a<b")
	} else {
		fmt.Println("a>b")
	}

	a.Add(a)
	fmt.Println(a)

	//结构体 初始化
	//rect1:=new(Rect)
	//rect2:=&Rect{}
	//rect3:=&Rect{0,0,100,200}
	//rect4:=&Rect{width:100,height:200}

	rect5 := NewRect(1, 1, 100, 200)

	fmt.Println("面积：", rect5.Area())
	//未显式进行初始化的变量都会被初始化为该类型的 ‘零值’
	// bool -> false  int -> 0  string -> ""

	foo := new(Foo)
	foo.Bar()
	foo.Foo()
}
