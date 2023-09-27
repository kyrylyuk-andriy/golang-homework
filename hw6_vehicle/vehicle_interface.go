package hw6_vehicle

type Vehicle interface {
	Stop()
	Start()
	ChangeSpeed(DeltaSpeed int)
	GetOutPassenger(amount int)
	GetInPassenger(amount int)
}
