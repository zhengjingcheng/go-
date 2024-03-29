package main

import "fmt"

//抽象层

// 水果类（抽象接口）
type Fruit interface {
	Show() //接口的某方法
}

// 工厂类 （抽象接口）
type AbstractFactory interface {
	CreateFruit() Fruit //生成水果类（抽象）的生产器方法
}

// 基础类模块
type Apple struct {
	Fruit //为了易于理解显示继承（此行可以省略）
}

func (apple *Apple) Show() {
	fmt.Println("我是苹果")
}

type Banana struct {
	Fruit //为了易于理解显示继承（此行可以省略）
}

func (apple *Banana) Show() {
	fmt.Println("我是香蕉")
}

type Pear struct {
	Fruit //为了易于理解显示继承（此行可以省略）
}

func (apple *Pear) Show() {
	fmt.Println("我是梨")
}

// 工厂模块
type AppleFactory struct {
	AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit
	//生产一个具体的苹果
	fruit = new(Apple)
	return fruit
}

type BananaFactory struct {
	AbstractFactory
}

func (fac *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit
	//生产一个具体的苹果
	fruit = new(Banana)
	return fruit
}

type PearFactory struct {
	AbstractFactory
}

func (fac *PearFactory) CreateFruit() Fruit {
	var fruit Fruit
	//生产一个具体的苹果
	fruit = new(Pear)
	return fruit
}

// 业务层逻辑
func main() {
	//需求1：需要一个具体的苹果对象
	//1-先要一个具体的苹果工厂
	var appleFac AbstractFactory
	appleFac = new(AppleFactory)
	//2-生产相对应的具体水果
	var apple Fruit
	apple = appleFac.CreateFruit()
	//3-调用相应的方法
	apple.Show()

	//需求2：需要一个具体的香蕉对象
	//1-先要一个具体的香蕉工厂
	var bananaFac AbstractFactory
	bananaFac = new(BananaFactory)
	//2-生产相对应的具体水果
	var banana Fruit
	banana = bananaFac.CreateFruit()
	//3-调用相应的方法
	banana.Show()
	//需求3：需要一个具体的梨对象
	//1-先要一个具体的梨工厂
	var pearFac AbstractFactory
	pearFac = new(PearFactory)
	//2-生产相对应的具体水果
	var pear Fruit
	pear = pearFac.CreateFruit()
	//3-调用相应的方法
	pear.Show()

}
