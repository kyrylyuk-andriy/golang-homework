package hw6_vehicle

type Car struct {
	Brand     string
	Speed     int
	EngineOn  bool
	Passenger int
}

func (c *Car) Stop() {
	c.EngineOn = false
}

func (c *Car) Start() {
	c.EngineOn = true
}

func (c *Car) ChangeSpeed(DeltaSpeed int) {
	c.Speed = c.Speed + DeltaSpeed
}

func (c *Car) GetOutPassenger(amount int) {
	c.Passenger = c.Passenger - amount
}

func (c *Car) GetInPassenger(amount int) {
	c.Passenger = c.Passenger + amount
}
