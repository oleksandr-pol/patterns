package behavior

import (
	"fmt"
	"sync"
)

type train interface {
	RequestArrival()
	Departure()
	PermitArrival()
}

type PassengerTrain struct {
	Mediator mediator
}

func (g *PassengerTrain) RequestArrival() {
	if g.Mediator.canLand(g) {
		fmt.Println("PassengerTrain: Landing")
	} else {
		fmt.Println("PassengerTrain: Waiting")
	}
}

func (g *PassengerTrain) Departure() {
	fmt.Println("PassengerTrain: Leaving")
	g.Mediator.notifyFree()
}

func (g *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival Permitted. Landing")
}

type GoodsTrain struct {
	Mediator mediator
}

func (g *GoodsTrain) RequestArrival() {
	if g.Mediator.canLand(g) {
		fmt.Println("GoodsTrain: Landing")
	} else {
		fmt.Println("GoodsTrain: Waiting")
	}
}

func (g *GoodsTrain) Departure() {
	g.Mediator.notifyFree()
	fmt.Println("GoodsTrain: Leaving")
}

func (g *GoodsTrain) PermitArrival() {
	fmt.Println("GoodsTrain: Arrival Permitted. Landing")
}

type mediator interface {
	canLand(train) bool
	notifyFree()
}

type stationManager struct {
	isPlatformFree bool
	lock           *sync.Mutex
	trainQueue     []train
}

func NewStationManger() *stationManager {
	return &stationManager{
		isPlatformFree: true,
		lock:           &sync.Mutex{},
	}
}

func (s *stationManager) canLand(t train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *stationManager) notifyFree() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.PermitArrival()
	}
}
