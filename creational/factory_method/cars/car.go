package cars

type ICar interface {
	GetMaxSpeed() int
	GetHP() int
}

type Car struct {
	maxSpeed int
	hp       int
}

func (c *Car) GetMaxSpeed() int {
	return c.maxSpeed
}

func (c *Car) GetHP() int {
	return c.hp
}
