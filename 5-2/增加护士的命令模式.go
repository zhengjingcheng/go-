package main

import "fmt"

//医生-命令接收者
type Doctor2 struct {
}

func (d *Doctor2) treatEye() {
	fmt.Println("医生治疗眼睛")
}
func (d *Doctor2) treatNose() {
	fmt.Println("医生治疗鼻子")
}

//抽象的命令
type Command interface {
	Treat()
}

//治疗眼睛的病单
type CommandTreateEye1 struct {
	doctor *Doctor2
}

func (cmd *CommandTreateEye1) Treat() {
	cmd.doctor.treatEye()
}

//治疗鼻子的病单
type CommandTreatNose1 struct {
	doctor *Doctor2
}

func (cmd *CommandTreatNose1) Treat() {
	cmd.doctor.treatNose()
}

//护士-调用命令者
type Nurse struct {
	CmdList []Command
}

//发送病单，发送命令的方法
func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}
	for _, cmd := range n.CmdList {
		cmd.Treat()
	}
}

//病人
func main() {
	//依赖病单，通过填写病单让医生看病
	doctor := new(Doctor2)
	//治疗眼睛的病单
	cmdEye := CommandTreateEye1{doctor}
	cmdNose := CommandTreatNose1{doctor}
	//护士
	nurse := new(Nurse)
	//收集管理病单
	nurse.CmdList = append(nurse.CmdList, &cmdEye)
	nurse.CmdList = append(nurse.CmdList, &cmdNose)
	nurse.Notify()
}



