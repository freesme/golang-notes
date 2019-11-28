package main

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
