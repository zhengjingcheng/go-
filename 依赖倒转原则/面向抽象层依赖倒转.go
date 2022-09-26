package main

import "fmt"

type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}
type BenZ struct {
}
type Bmw struct {
}

func (benz *BenZ) Run() {
	fmt.Println("Benz is runing")
}
func (benz *Bmw) Run() {
	fmt.Println("Benz is runing")
}

type Li_4 struct {
	//...
}

func (li4 *Li_4) Drive(car Car) {
	fmt.Println("li4 drive car")
	car.Run()
}

type Zhang_3 struct {
	//...
}

func (zhang3 *Zhang_3) Drive(car Car) {
	fmt.Println("Zhang3 drive car")
	car.Run()
}
func main() {
	var bmw Car
	var benci Car
	bmw = &Bmw{}
	benci = &BenZ{}
	var zhang3 Driver
	zhang3 = &Zhang_3{}
	zhang3.Drive(bmw)
	var li4 Driver
	li4 = &Li_4{}
	li4.Drive(benci)
}
