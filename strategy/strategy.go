package strategy

import "time"

type Interface interface {
  CalculateCost(distance float64) float64
  CalculateETA(startAt time.Time, td TripDetail) time.Time
}

type Location struct {
  CityName string
  Lat, Lng float64
}

type TripDetail struct {
  Origin, Destination Location
}

func (t TripDetail) Distance() float64 {
  // calculate distance from origin to destination
  return 10
}
