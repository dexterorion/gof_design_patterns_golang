package eletronics

type ILaptop interface {
	GetMemory() int
	GetCPU() string
}

type Laptop struct {
	memory int
	cpu    string
}

func (l *Laptop) GetMemory() int {
	return l.memory
}

func (l *Laptop) GetCPU() string {
	return l.cpu
}
