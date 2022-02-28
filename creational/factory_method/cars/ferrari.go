package cars

type Ferrari struct {
	Car
}

func NewFerrari() ICar {
	return &Ferrari{
		Car: Car{
			maxSpeed: 400,
			hp:       1000,
		},
	}
}
