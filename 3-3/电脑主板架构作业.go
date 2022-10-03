package main

import "fmt"

// 抽象类
// 抽象显卡
type GpuAbst interface {
	ShowGpu()
}

// 抽象内存
type RamAbst interface {
	ShowRam()
}

// 抽象CPU
type CpuAbst interface {
	ShowCpu()
}

// 抽象工厂
type FactAbst interface {
	CreateGpu() GpuAbst
	CreateRam() RamAbst
	CreateCpu() CpuAbst
}

// 实现层
// Inter显卡
type InterGpu struct {
}

func (Ig *InterGpu) ShowGpu() {
	fmt.Println("inter显卡")
}

// Inter内存
type InterRam struct {
}

func (Ig *InterRam) ShowRam() {
	fmt.Println("inter内存")
}

// InterCPU
type InterCpu struct {
}

func (Ig *InterCpu) ShowCpu() {
	fmt.Println("interCPU")
}

type InterFact struct {
}

func (IF *InterFact) CreateGpu() GpuAbst {
	var fa GpuAbst
	fa = new(InterGpu)
	return fa
}
func (IF *InterFact) CreateRam() RamAbst {
	var fa RamAbst
	fa = new(InterRam)
	return fa
}
func (IF *InterFact) CreateCpu() CpuAbst {
	var fa CpuAbst
	fa = new(InterCpu)
	return fa
}

// Nvidia显卡
type NvidiaGpu struct {
}

func (Ig *NvidiaGpu) ShowGpu() {
	fmt.Println("Nvidia显卡")
}

// Nvidia内存
type NvidiaRam struct {
}

func (Ig *NvidiaRam) ShowRam() {
	fmt.Println("inter内存")
}

// NvidiaCPU
type NvidiaCpu struct {
}

func (Ig *NvidiaCpu) ShowCpu() {
	fmt.Println("NvidiaCPU")
}

type NvidiaFact struct {
}

func (IF *NvidiaFact) CreateGpu() GpuAbst {
	var fa GpuAbst
	fa = new(NvidiaGpu)
	return fa
}
func (IF *NvidiaFact) CreateRam() RamAbst {
	var fa RamAbst
	fa = new(NvidiaRam)
	return fa
}
func (IF *NvidiaFact) CreateCpu() CpuAbst {
	var fa CpuAbst
	fa = new(NvidiaCpu)
	return fa
}

// Kingston显卡
type KingstonGpu struct {
}

func (Ig *KingstonGpu) ShowGpu() {
	fmt.Println("Kingston显卡")
}

// Kingston内存
type KingstonRam struct {
}

func (Ig *KingstonRam) ShowRam() {
	fmt.Println("Kingston内存")
}

// KingstonCPU
type KingstonCpu struct {
}

func (Ig *KingstonCpu) ShowCpu() {
	fmt.Println("KingstonCPU")
}

type KingstonFact struct {
}

func (IF *KingstonFact) CreateGpu() GpuAbst {
	var fa GpuAbst
	fa = new(KingstonGpu)
	return fa
}
func (IF *KingstonFact) CreateRam() RamAbst {
	var fa RamAbst
	fa = new(KingstonRam)
	return fa
}
func (IF *KingstonFact) CreateCpu() CpuAbst {
	var fa CpuAbst
	fa = new(KingstonCpu)
	return fa
}

func main() {
	// 1台（Intel的CPU，Intel的显卡，Intel的内存）
	var Interf FactAbst
	Interf = new(InterFact)
	cpu1 := Interf.CreateCpu()
	cpu1.ShowCpu()
	ram1 := Interf.CreateRam()
	ram1.ShowRam()
	gpu1 := Interf.CreateGpu()
	gpu1.ShowGpu()
	//1台（Intel的CPU， nvidia的显卡，Kingston的内存）
	var nvf FactAbst
	nvf = new(NvidiaFact)
	cpu2 := Interf.CreateCpu()
	cpu2.ShowCpu()
	gpu2 := nvf.CreateGpu()
	gpu2.ShowGpu()
	var kgf FactAbst
	kgf = new(KingstonFact)
	ram2 := kgf.CreateRam()
	ram2.ShowRam()
}
