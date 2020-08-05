package main

import "fmt"

//抽象层
type Card interface {
	Display()
}

type Memory interface {
	Storage()
}

type CPU interface {
	Calculate()
}

type Computer struct {
	cpu CPU
	mem Memory
	card Card
}

func NewComputer(cpu CPU,mem Memory,card Card) *Computer{
	return &Computer{
		cpu: cpu,
		mem: mem,
		card: card,
	}
}

func (computer *Computer) Dowork(){
	computer.cpu.Calculate()
	computer.mem.Storage()
	computer.card.Display()
}

//实现层
type IntelCPU struct {
	CPU
}
func (intel *IntelCPU) Calculate(){
	fmt.Println("Intel Cpu开始计算了")
}

type IntelMemory struct {
	Memory
}

func (intelM *IntelMemory) Storage(){
	fmt.Println("Intel Memory 开始存储了")
}

type IntelCard struct {
	Card
}

func (interC *IntelCard) Display(){
	fmt.Println("Intel Card 开始显示了")
}

//kinston
type KingstonMemory struct {
	Memory
}

func (kins *KingstonMemory) Storage(){
	fmt.Println("kinstone memory storage")
}

//nvidia
type NvidiaCard struct {
	Card
}

func (nvidia *NvidiaCard) Display(){
	fmt.Println("Nvidia card display")
}

