package hw6_vehicle

import "fmt"

type Line struct {
	LineNumber    int
	DepartureCity string
	ArrivalCity   string
	vehiclesList  []Vehicle
}

func (l *Line) AddVehicleToLine(vehicle Vehicle) {
	l.vehiclesList = append(l.vehiclesList, vehicle)
}

func (l Line) ShowPublicTransportationStatus() {
	fmt.Printf("Show list of public transports on a line %d from %s to %s\n", l.LineNumber, l.DepartureCity, l.ArrivalCity)
	fmt.Println(l.vehiclesList)
}
