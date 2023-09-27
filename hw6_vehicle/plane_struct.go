package hw6_vehicle

type Plane struct {
	Brand     string
	Speed     int
	EngineOn  bool
	Passenger int
}

func (p *Plane) Stop() {
	p.EngineOn = false
}

func (p *Plane) Start() {
	p.EngineOn = true
}

func (p *Plane) ChangeSpeed(DeltaSpeed int) {
	p.Speed = p.Speed + DeltaSpeed
}

func (p *Plane) GetOutPassenger(amount int) {
	p.Passenger = p.Passenger - amount
}

func (p *Plane) GetInPassenger(amount int) {
	p.Passenger = p.Passenger + amount
}
