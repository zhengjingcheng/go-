package main

import (
	"fmt"
	"sync"
)

//由于加锁操作很浪费时间我们可以利用go语言自己的Once模块实现单例模式

var once sync.Once

type singelton4 struct {
}

var instance4 *singelton4

func GetInstance4() *singelton4 {
	once.Do(func() {
		if instance4 == nil {
			instance4 = new(singelton4)
		}
	})
	return instance4
}
func (s *singelton4) SomeThing() {
	fmt.Println("单例对象的某方法")
}
func main() {
	s := GetInstance4()
	s.SomeThing()
}