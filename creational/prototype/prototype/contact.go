package prototype

import "fmt"

func NewContact(name, phone string) Node {
	return &contact{
		name:  name,
		phone: phone,
	}
}

type contact struct {
	name  string
	phone string
}

func (c *contact) Print(indentation string) {
	fmt.Println(indentation + c.name + " - " + c.phone)
}

func (c *contact) Clone() Node {
	return NewContact(c.name+"_clone", c.phone)
}
