package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type MP3 interface {
	PlayMusic()
}

type Phone struct {
}

//手机实现了USB接口 和MP3接口
func (p Phone) Start() {
	fmt.Println("USB 接入手机")
}
func (p Phone) Stop() {
	fmt.Println("USB 拔出手机")
}
func (p Phone) PlayMusic() {
	fmt.Println("手机播放MPS")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("USB 插入相机")
}
func (c Camera) Stop() {
	fmt.Println("USB 拔出相机")
}

type Compture struct {
}

func (compture Compture) Use(u Usb) {
	u.Start()
}
func (compture Compture) Out(u Usb) {
	u.Stop()
}

func (compture Compture) PlayMusic(play MP3) {
	play.PlayMusic()
}

func main() {

	phone := new(Phone)
	camera := new(Camera)

	var compture Compture
	//Usb接口可以接收实现类的实例
	compture.Use(phone)
	compture.Use(camera)
	compture.Out(phone)
	compture.Out(camera)

	compture.PlayMusic(phone) //不太恰当

}
