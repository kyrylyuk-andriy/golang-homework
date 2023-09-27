package vehicle

type Venicle interface {
	Stop()
	Start()
	ChangeSpeed() int
	GetOutPassenger()
	GetInPassenger()
}

type Car struct {
	Brand    string
	Speed    int
	EngineOn bool
}

type Train struct {
	Brand    string
	Speed    int
	EngineOn bool
}

type Plane struct {
	Brand    string
	Speed    int
	EngineOn bool
}
