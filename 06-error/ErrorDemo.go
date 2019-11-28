package main

import (
	"errors"
	"fmt"
)

//自定义 error

type PathError struct {
	Op   string
	Path string
	Err  error
}

// 实现Error的方法 编译器即认定此实现了Error接口
func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

//自定义类型错误
/*
	errors.New("错误说明") 会返回一个error类型的值，表示一个错误
	panic内置函数，接收一个interface{}类型的值（也就是任何值了）作为参数，可以接收error类习惯的变量，输出错误信息，并退出程序
*/

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误。。。")
	}
}

func main() {
	//err:=readConf("config.ini")
	err := readConf("config2.ini")
	if err != nil {
		//终止程序
		panic(err)
	}
	fmt.Println("next code running...")
}
