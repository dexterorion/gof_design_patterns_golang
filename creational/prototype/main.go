package main

import (
	"fmt"

	"github.com/dexterorion/gof_design_patterns_golang/creational/prototype/prototype"
)

func main() {
	contact1 := prototype.NewContact("Contact1", "1111")
	contact2 := prototype.NewContact("Contact2", "2222")
	contact3 := prototype.NewContact("Contact3", "3333")
	agenda1 := prototype.NewAgenda("Sub-Agenda1", []prototype.Node{contact1})
	agenda2 := prototype.NewAgenda("Main-Agenda", []prototype.Node{agenda1, contact2, contact3})

	fmt.Println("Printing main agenda with its sub-agendas")
	agenda2.Print("\t")

	clonedAgenda := agenda2.Clone()

	fmt.Println("Printing cloned agenda")
	clonedAgenda.Print("\t")
}
