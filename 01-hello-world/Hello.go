//注释
/*
	长注释
*/
package main //导入包 -main函数 在 main包下才可以运行

//导入包
//import "fmt"  //导入单个包 	//格式化输出的包 Ctrl+鼠标左键点击fmt 进入查看详细定义 官网文档：https://golang.org/pkg/fmt/
import ( //批量导入包
	"fmt"
	//	"time"  //go语言中不允许引用未使用的包 否则编译器报编译错误
)

//函数定义 关键字 func
func main() { //强制 左花括号
	fmt.Println("Hello Golang") //打印 Hello Golang

	//接收多返回值
	result1, err := Compute(1, 2)
	//错误处理
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("result:", result1)
	fmt.Println("------------------------------")
	//问题追踪调试
	//定义变量 使用  :=  初始化
	favl := 110.46
	ival := 100
	sval := "Kakakak"
	fmt.Println("This is ", sval) //字符串拼接
	fmt.Printf("favl=%f,ival=%d,sval=%s\n", favl, ival, sval)
	fmt.Printf("favl=%v,ival=%v,sval=%v\n", favl, ival, sval)
}

//函数定义 支持多返回值
/*
func 函数名(参数列表)(返回值列表){
	//函数体
}
*/
func Compute(value1 int, value2 float64) (result1 float64, err error) {
	return value2, nil
}
