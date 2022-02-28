package eletronics

type IGalaxy interface{}

func NewGalaxy(mobile Mobile) IGalaxy {
	return &Galaxy{
		Mobile: mobile,
	}
}

type Galaxy struct {
	Mobile
}
