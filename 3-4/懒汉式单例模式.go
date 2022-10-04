package main

import "fmt"

//上面代码推演了一个单例的创建和逻辑过程，上述是单例模式中的一种，属于“饿汉式”。
//含义是，在初始化单例唯一指针的时候，就已经提前开辟好了一个对象，申请了内存。
//饿汉式的好处是，不会出现线程并发创建，导致多个单例的出现，
//但是缺点是如果这个单例对象在业务逻辑没有被使用，也会客观的创建一块内存对象。
//那么与之对应的模式叫“懒汉式”，代码如下：
type singelton2 struct {
}

var instance2 *singelton2

func GetInstance2() *singelton2 {
	if instance2 == nil {
		instance2 = new(singelton2)
	}
	return instance2
}
func (s *singelton2) SomeThing() {
	fmt.Println("单例对象的某方法")
}
func main() {
	s := GetInstance2()
	s.SomeThing()
}