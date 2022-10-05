package main

import "fmt"

type Cooker struct {
}

func (c *Cooker) MakeChicken() {
	fmt.Println("烤串师傅烤了鸡肉串")
}
func (c *Cooker) MakeChuaner() {
	fmt.Println("烤肉师傅烤了羊肉串")
}

//抽象的命令
type Command interface {
	Make()
}
type CommandCookChicken struct {
	cooker *Cooker
}

func (cmd *CommandCookChicken) Make() {
	cmd.cooker.MakeChicken()
}

type CommandCookChuaner struct {
	cooker *Cooker
}

func (cmd *CommandCookChuaner) Make() {
	cmd.cooker.MakeChuaner()
}

type WaiterMM struct {
	CmdList []Command
}

func (w *WaiterMM) Notify() {
	if w.CmdList == nil {
		return
	}
	for _, cmd := range w.CmdList {
		cmd.Make()
	}
}
func main() {
	cooker := new(Cooker)
	cmdChicken := CommandCookChicken{cooker}
	cmdChuaner := CommandCookChuaner{cooker}
	mm := new(WaiterMM)
	mm.CmdList = append(mm.CmdList, &cmdChicken)
	mm.CmdList = append(mm.CmdList, &cmdChuaner)
	mm.Notify()
}
