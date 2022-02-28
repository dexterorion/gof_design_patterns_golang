package cars

type McLaren struct {
	Car
}

func NewMcLaren() ICar {
	return &McLaren{
		Car: Car{
			maxSpeed: 440,
			hp:       1200,
		},
	}
}
