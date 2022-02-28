package eletronics

import "errors"

type IFactory interface {
	MakeLaptop() ILaptop
	MakeMobile() IMobile
}

func GetEletronicsFactory(brand string) (IFactory, error) {
	if brand == "apple" {
		return &Apple{}, nil
	} else {
		if brand == "samsung" {
			return &Samsung{}, nil
		}
	}

	return nil, errors.New("brand is invalid")
}
