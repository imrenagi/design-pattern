package strategy

import "time"

type Road struct {
}

func (r Road) CalculateCost(distance float64) float64 {
  return 100000 + distance * 5000
}

const (
  defaultSpeed = 80 // kmh
)

func (r Road) CalculateETA(startAt time.Time, trip TripDetail) time.Time {
  t := trip.Distance() / defaultSpeed
  var y int = int(t)
  duration := time.Duration(y) * time.Hour
  return startAt.Add(duration)
}

