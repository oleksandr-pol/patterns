package main

import (
	"../../behavior"
)

func main() {
	stationManager := behavior.NewStationManger()
	passengerTrain := &behavior.PassengerTrain{
		Mediator: stationManager,
	}
	goodsTrain := &behavior.GoodsTrain{
		Mediator: stationManager,
	}
	passengerTrain.RequestArrival()
	goodsTrain.RequestArrival()
	passengerTrain.Departure()
}
