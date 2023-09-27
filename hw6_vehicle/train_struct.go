package hw6_vehicle

type Train struct {
	Brand     string
	Speed     int
	EngineOn  bool
	Passenger int
}

func (t *Train) Stop() {
	t.EngineOn = false
}

func (t *Train) Start() {
	t.EngineOn = true
}

func (t *Train) ChangeSpeed(DeltaSpeed int) {
	t.Speed = t.Speed + DeltaSpeed
}

func (t *Train) GetOutPassenger(amount int) {
	t.Passenger = t.Passenger - amount
}

func (t *Train) GetInPassenger(amount int) {
	t.Passenger = t.Passenger + amount
}
