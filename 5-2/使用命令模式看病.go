package main

import "fmt"

type Doctor1 struct {
}

func (d *Doctor1) treatEye() {
	fmt.Println("医生治疗眼睛")
}
func (d *Doctor1) treatNose() {
	fmt.Println("医生治疗鼻子")
}

// 治疗眼睛的病单
type CommandTreateEye struct {
	doctor *Doctor1
}

func (cmd *CommandTreateEye) Treat() {
	cmd.doctor.treatEye()
}

// 治疗鼻子的病单
type CommandTreatNose struct {
	doctor *Doctor1
}

func (cmd *CommandTreatNose) Treat() {
	cmd.doctor.treatNose()
}

// 病人
func main() {
	doctor := new(Doctor1)
	cmdEye := CommandTreateEye{doctor}
	cmdEye.Treat()
	cmdNose := CommandTreatNose{doctor: doctor}
	cmdNose.Treat()
}
