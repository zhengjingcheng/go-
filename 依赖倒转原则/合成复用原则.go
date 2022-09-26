package main

import "fmt"

type Cat struct {
}

func (c *Cat) Eat() {
	fmt.Println("小猫吃饭")
}

//给小猫添加一个可以睡觉的方法

type CatB struct {
	Cat
}

func (cb *CatB) sleep() {
	fmt.Println("小猫睡觉")
}

// 使用组合的方法给小猫添加一个可以睡觉的方法
type Catc struct {
	C *Cat
}

func (cc *Catc) sleep() {
	fmt.Println("小猫睡觉")
}

func main() {
	//通过继承的方式
	cb := new(CatB)
	cb.Eat()
	cb.sleep()
	cc := new(Catc)
	cc.sleep()
	cc.C.Eat()
}
