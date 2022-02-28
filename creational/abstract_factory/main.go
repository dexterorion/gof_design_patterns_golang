package main

import (
	"fmt"

	"github.com/dexterorion/gof_design_patterns_golang/creational/abstract_factory/eletronics"
)

func main() {
	appleFactory, _ := eletronics.GetEletronicsFactory("apple")
	samsungFactory, _ := eletronics.GetEletronicsFactory("samsung")

	iphone := appleFactory.MakeMobile()
	macbook := appleFactory.MakeLaptop()

	// print apple products
	fmt.Printf("Printing apple products\n")
	fmt.Printf("Iphone memory is [%dgb] and has [%d] SIM cards space\n", iphone.GetMemory(), iphone.GetChipSlotsQuantity())
	fmt.Printf("Macbook memory is [%dgb] and has [%s] CPU\n", macbook.GetMemory(), macbook.GetCPU())

	fmt.Printf("\n\n\n")

	galaxy := samsungFactory.MakeMobile()
	odyssey := samsungFactory.MakeLaptop()

	fmt.Printf("Printing samsung products\n")
	fmt.Printf("Samsung galaxy memory is [%dgb] and has [%d] SIM cards space\n", galaxy.GetMemory(), galaxy.GetChipSlotsQuantity())
	fmt.Printf("Samsung odyssey memory is [%dgb] and has [%s] CPU\n", odyssey.GetMemory(), odyssey.GetCPU())
}
