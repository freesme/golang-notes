package main

import "os"

type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}

type IReader interface {
	Read(buf []byte) (n int, err error)
}
type IWriter interface {
	Write(buf []byte) (n int, err error)
}

type ICloser interface {
	Close() error
}

//尽管File类并没有从这些接口继承，甚至可以不知道这些接口的存在，但是File类实现了这些接口，可以进行赋值

//----- 接口赋值-----

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type Lesser interface {
	Less(b Integer) bool
}

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Add(b Integer) {
	*a += b
}

//2.将一个接口赋值给另一个接口

func main() {

	//var ifile IFile= new(os.File)
	//
	//var a Integer = 1
	////1.将对象实例赋值给接口
	//var b LessAdder = &a
	//
	//var b1 Lesser = &a
	//var b2 Lesser = a

}
