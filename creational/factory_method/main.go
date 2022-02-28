package main

import (
	"fmt"

	"github.com/dexterorion/gof_design_patterns_golang/creational/factory_method/cars"
)

func main() {
	ferrari, _ := cars.GetCar("ferrari")
	maclaren, _ := cars.GetCar("mclaren")

	// print ferrari
	fmt.Printf("Printing ferrari\n")
	fmt.Printf("Ferrari max speed is [%dkm/h] and has [%dhp]\n", ferrari.GetMaxSpeed(), ferrari.GetHP())

	fmt.Printf("\n\n\n")

	// print ferrari
	fmt.Printf("Printing maclaren\n")
	fmt.Printf("Mclaren max speed is [%dkm/h] and has [%dhp]\n", maclaren.GetMaxSpeed(), maclaren.GetHP())

}
