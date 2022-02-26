package prototype

import "fmt"

func NewAgenda(name string, children []Node) Node {
	return &agenda{
		name:     name,
		children: children,
	}
}

type agenda struct {
	children []Node
	name     string
}

func (a *agenda) Print(indentation string) {
	fmt.Println(indentation + a.name)
	for _, i := range a.children {
		i.Print(indentation + indentation)
	}
}

func (a *agenda) Clone() Node {
	var tempChildren []Node
	for _, i := range a.children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}
	return NewAgenda(a.name+"_clone", tempChildren)
}
