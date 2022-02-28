package cars

import "errors"

func GetCar(carType string) (ICar, error) {
	if carType == "ferrari" {
		return NewFerrari(), nil
	} else {
		if carType == "mclaren" {
			return NewMcLaren(), nil
		}
	}
	return nil, errors.New("car type is invalid")
}
