package main

import "fmt"

// 抽象层
// 抽象的构件
type Phone interface {
	Show() //构件的功能
}

// 装饰器基础类
type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {}

// 实现层
// 具体的构件
type Huawei struct {
}

func (hw *Huawei) Show() {
	fmt.Println("秀出了华为手机")
}

type Xiaomi struct {
}

func (xm *Xiaomi) Show() {
	fmt.Println("秀出了小米手机")
}

// 具体的装饰器类
type MoDecorator struct {
	Decorator //继承基础装饰器类
}

func (md *MoDecorator) Show() {
	md.phone.Show()
	fmt.Println("贴膜的手机")
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

type KeDecorator struct {
	Decorator
}

func (md *KeDecorator) Show() {
	md.phone.Show()
	fmt.Println("手机壳的手机")
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone}}
}

// 业务逻辑层
func main() {
	var huawei Phone
	huawei = new(Huawei)
	huawei.Show() //调用原构件方法

	fmt.Println("---------")
	//用贴膜装饰器装饰得到新功能构件
	var moHuawei Phone
	moHuawei = NewMoDecorator(huawei)
	moHuawei.Show()
	fmt.Println("---------")
	//用贴膜装饰器装饰得到新功能构件
	var keHuawei Phone
	keHuawei = NewKeDecorator(huawei)
	keHuawei.Show()
	fmt.Println("---------")
	//用贴膜装饰器装饰得到新功能构件
	var kemoHuawei Phone
	kemoHuawei = NewMoDecorator(keHuawei)
	kemoHuawei.Show()
}
