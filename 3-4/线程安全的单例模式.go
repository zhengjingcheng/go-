package main

import (
	"fmt"
	"sync"
)

//上面的“懒汉式”实现是非线程安全的设计方式，
//也就是如果多个线程或者协程同时首次调用GetInstance()方法有概率导致多个实例被创建，则违背了单例的设计初衷。
//那么在上面的基础上进行修改，可以利用Sync.Mutex进行加锁，保证线程安全。
//这种线程安全的写法，有个最大的缺点就是每次调用该方法时都需要进行锁操作，在性能上相对不高效，具体的实现改进如下

//定义锁
var lock sync.Mutex

type singelton3 struct {
}

var instance3 *singelton3

func GetInstance3() *singelton3 {
	//为了线程安全，增加互斥
	lock.Lock()
	defer lock.Unlock()
	if instance3 == nil {
		instance3 = new(singelton3)
	}
	return instance3
}

func (s *singelton3) SomeThing() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstance3()
	s.SomeThing()
}
