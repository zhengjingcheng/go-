package main

import "fmt"

//	简单工厂模式并不属于GoF的23种设计模式。他是开发者自发认为的一种非常简易的设计模式，其角色和职责如下：
//	工厂（Factory）角色：简单工厂模式的核心，它负责实现创建所有实例的内部逻辑。工厂类可以被外界直接调用，创建所需的产品对象。
//	抽象产品（AbstractProduct）角色：简单工厂模式所创建的所有对象的父类，它负责描述所有实例所共有的公共接口。
//	具体产品（Concrete Product）角色：简单工厂模式所创建的具体实例对象。

// ======= 抽象层 =========
//水果类(抽象接口)

type Fruit2 interface {
	Show() //接口的某方法
}

// ======= 基础类模块 =========
type Apple struct {
	Fruit2 //为了易于理解显示继承（此行可以省略）
}

func (apple *Apple) Show() {
	fmt.Println("我是苹果")
}

type Banana struct {
	Fruit2
}

func (banana *Banana) Show() {
	fmt.Println("我是香蕉")
}

type Pear struct {
	Fruit2
}

func (banana *Pear) Show() {
	fmt.Println("我是梨")
}

// ========= 工厂模块  =========
//一个工厂， 有一个生产水果的机器，返回一个抽象水果的指针

type Factory struct {
}

func (fac *Factory) CreaterFruit(kind string) Fruit2 {
	var fruit Fruit2
	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}
	return fruit
}

// ==========业务逻辑层==============
func main() {
	factory := new(Factory)

	apple := factory.CreaterFruit("apple")
	apple.Show()

	banana := factory.CreaterFruit("banana")
	banana.Show()

	pear := factory.CreaterFruit("pear")
	pear.Show()
}
