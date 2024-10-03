package main

type Mediator interface {
	canArrive(t Train) bool
	notifyAboutDeparture()
}

type StationManager struct {
	isPlatformFree bool
	trainQueue     []Train
}

func newStationManger() *StationManager {
	return &StationManager{
		isPlatformFree: true,
	}
}

func (s *StationManager) canArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) notifyAboutDeparture() {
	s.isPlatformFree = true
	if len(s.trainQueue) > 0 {
		s.isPlatformFree = false
		firstTrain := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrain.permitArrival()
	}
}
