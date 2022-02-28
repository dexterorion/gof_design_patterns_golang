package eletronics

type Samsung struct{}

func (s *Samsung) MakeLaptop() ILaptop {
	return &Macbook{
		Laptop: Laptop{
			memory: 32,
			cpu:    "Intel i9",
		},
	}
}

func (s *Samsung) MakeMobile() IMobile {
	return &Iphone{
		Mobile: Mobile{
			memory:       32,
			chipSlotsQty: 4,
		},
	}
}
