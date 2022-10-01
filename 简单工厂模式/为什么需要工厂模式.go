package main

import "fmt"

//首先来看，如果没有工厂模式，在开发者创建一个类的对象时，如果有很多不同种类的对象将会如何实现，代码如下：

// 水果类
type Fruit struct {
	// ...
	// ...
	// ...
}

func (f *Fruit) Show(name string) {
	if name == "apple" {
		fmt.Println("我是苹果")
	} else if name == "banana" {
		fmt.Println("我是香蕉")
	} else if name == "pear" {
		fmt.Println("我是梨")
	}
}

// 创建一个Fruit对象
func NewFruit(name string) *Fruit {
	fruit := new(Fruit)
	if name == "apple" {
		//创建apple逻辑
	} else if name == "banana" {
		//创建banana逻辑
	} else if name == "pear" {
		//创建梨逻辑
	}
	return fruit
}

func main() {
	apple := NewFruit("apple")
	apple.Show("apple")
	banana := NewFruit("banana")
	banana.Show("banana")

	pear := NewFruit("pear")
	pear.Show("pear")
}
