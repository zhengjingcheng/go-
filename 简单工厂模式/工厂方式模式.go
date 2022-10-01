package main

import "fmt"

//抽象工厂（Abstract Factory）角色：工厂方法模式的核心，任何工厂类都必须实现这个接口。
//工厂（Concrete Factory）角色：具体工厂类是抽象工厂的一个实现，负责实例化产品对象。
//抽象产品（Abstract Product）角色：工厂方法模式所创建的所有对象的父类，它负责描述所有实例所共有的公共接口。
//具体产品（Concrete Product）角色：工厂方法模式所创建的具体实例对象。

// 水果类(抽象接口)
type Fruit3 interface {
	Show() //接口的显示方法
}

// 工厂类（抽象接口）
type AbstractFactory interface {
	CreateFruit() Fruit //生产水果类（抽象）的生产器方法
}

// 基础类模块
type Apple3 struct {
	Fruit3 //为了易于理解显示继承（此行可以省略）
}

func (apple *Apple3) Show() {
	fmt.Println("我是苹果")
}

type Banana3 struct {
	Fruit3 //为了易于理解显示继承（此行可以省略）
}

func (apple *Banana3) Show() {
	fmt.Println("我是香蕉")
}

type Pear3 struct {
	Fruit3 //为了易于理解显示继承（此行可以省略）
}

func (apple *Pear3) Show() {
	fmt.Println("我是梨")
}
