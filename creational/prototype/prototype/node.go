package prototype

type Node interface {
	Print(indentation string)
	Clone() Node
}
