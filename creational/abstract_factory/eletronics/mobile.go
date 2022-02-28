package eletronics

type IMobile interface {
	GetMemory() int
	GetChipSlotsQuantity() int
}

type Mobile struct {
	memory       int
	chipSlotsQty int
}

func (m *Mobile) GetMemory() int {
	return m.memory
}

func (m *Mobile) GetChipSlotsQuantity() int {
	return m.chipSlotsQty
}
