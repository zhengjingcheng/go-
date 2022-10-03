package main

import (
	"fmt"
)

// 抽象层
type AbstractApple interface {
	ShowApple()
}
type AbstractBanana interface {
	ShowBanana()
}
type AbstractPear interface {
	ShowPear()
}

// 抽象工厂
type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

// 实现层
// 中国产品族
type ChinApple struct {
}

func (ca *ChinApple) ShowApple() {
	fmt.Println("中国苹果")
}

type ChinaBanana struct {
}

func (cb *ChinaBanana) ShowBanana() {
	fmt.Println("中国香蕉")
}

type ChinaPear struct {
}

func (cb *ChinaPear) ShowPear() {
	fmt.Println("中国梨")
}

type ChinaFactory struct {
}

func (cf *ChinaFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(ChinApple)
	return apple
}
func (cf *ChinaFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(ChinaBanana)
	return banana
}
func (cf *ChinaFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(ChinaPear)
	return pear
}

// 美国产品族
type UsaApple struct {
}

func (ca *UsaApple) ShowApple() {
	fmt.Println("美国苹果")
}

type UsaBanana struct {
}

func (cb *UsaBanana) ShowBanana() {
	fmt.Println("美国香蕉")
}

type UsaPear struct {
}

func (cb *UsaPear) ShowPear() {
	fmt.Println("美国梨")
}

type UsaFactory struct {
}

func (cf *UsaFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(UsaApple)
	return apple
}
func (cf *UsaFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(UsaBanana)
	return banana
}
func (cf *UsaFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(UsaPear)
	return pear
}
func main() {
	//需求1: 需要美国的苹果、香蕉、梨 等对象
	//1-创建一个美国工厂
	var aFac AbstractFactory
	aFac = new(UsaFactory)
	//2-生产美国苹果
	var aApple AbstractApple
	aApple = aFac.CreateApple()
	aApple.ShowApple()
	//美国梨
	var aPeer AbstractPear
	aPeer = aFac.CreatePear()
	aPeer.ShowPear()
	//美国香蕉
	var aBanana AbstractBanana
	aBanana = aFac.CreateBanana()
	aBanana.ShowBanana()
	//需求1: 需要中国的香蕉
	var cFac AbstractFactory
	cFac = new(ChinaFactory)

	var cbanana AbstractBanana
	cbanana = cFac.CreateBanana()
	cbanana.ShowBanana()

}
