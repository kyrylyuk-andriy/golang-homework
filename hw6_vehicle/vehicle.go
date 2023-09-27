package hw6_vehicle

func PublicTransportation() {

	// Lets add three vehicles, each one per car, train and plane, all vehicles are in idle state
	car1 := Car{Brand: "VW", Speed: 0, EngineOn: false, Passenger: 0}
	train1 := Train{Brand: "GM", Speed: 0, EngineOn: false, Passenger: 0}
	plane1 := Plane{Brand: "Boeing", Speed: 0, EngineOn: false, Passenger: 0}

	// Initialize Line (Route)
	line1 := Line{LineNumber: 1, DepartureCity: "Kyiv", ArrivalCity: "London"}

	car1.Start()
	car1.GetInPassenger(1)
	car1.ChangeSpeed(50)
	line1.AddVehicleToLine(&car1)

	train1.Start()
	train1.GetInPassenger(100)
	train1.ChangeSpeed(500)
	line1.AddVehicleToLine(&train1)

	plane1.Start()
	plane1.GetInPassenger(300)
	plane1.ChangeSpeed(1000)
	line1.AddVehicleToLine(&plane1)

	line1.ShowPublicTransportationStatus()

}
