package main

import "fmt"

// 适配目标
type Attack interface {
	Fight()
}

// 具体的技能
type Dabaojian struct {
}

func (da *Dabaojian) Fight() {
	fmt.Println("使用了大保健技能，将敌人击杀")
}

type Hero struct {
	Name   string
	attack Attack
}

func (h *Hero) Skill() {
	fmt.Println(h.Name, "使用了技能")
	h.attack.Fight() //使用具体的战斗方式
}

// 适配者
type PowerOff struct {
}

func (p *PowerOff) ShutDown() {
	fmt.Println("电脑即将关机")
}

// 适配器
type Adapter1 struct {
	poweroff *PowerOff
}

func (a *Adapter1) Fight() {
	a.poweroff.ShutDown()
}
func NewAdapter1(p *PowerOff) *Adapter1 {
	return &Adapter1{p}
}

func main() {
	gailun := Hero{Name: "盖伦", attack: NewAdapter1(new(PowerOff))}
	gailun.Skill()
}
