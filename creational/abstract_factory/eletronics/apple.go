package eletronics

type Apple struct{}

func (a *Apple) MakeLaptop() ILaptop {
	return &Macbook{
		Laptop: Laptop{
			memory: 16,
			cpu:    "M1",
		},
	}
}

func (a *Apple) MakeMobile() IMobile {
	return &Iphone{
		Mobile: Mobile{
			memory:       16,
			chipSlotsQty: 2,
		},
	}
}
