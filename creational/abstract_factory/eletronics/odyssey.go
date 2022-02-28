package eletronics

type IOdyssey interface{}

func NewOdyssey(laptop Laptop) IOdyssey {
	return &Odyssey{
		Laptop: laptop,
	}
}

type Odyssey struct {
	Laptop
}
