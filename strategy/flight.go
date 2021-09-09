package strategy

import "time"

type WeatherPredictor interface {
  Predict(cityName string) bool
}

type WPImpl struct {
}

func (wp WPImpl) Predict(cityName string) bool {
  return false
}

type Flight struct {
  wp WeatherPredictor
}

func (f Flight) CalculateCost(distance float64) float64 {
  return (10000000 + distance * 500000)/200
}

const (
  defaultAirSpeed = 500 // kmh
)


func (f Flight) CalculateETA(startAt time.Time, trip TripDetail) time.Time {
  t := trip.Distance() / defaultSpeed
  var y int = int(t)
  duration := time.Duration(y) * time.Hour

  eta := startAt.Add(duration)
  ok := f.wp.Predict(trip.Origin.CityName)
  if !ok {
    eta = eta.Add(1 * time.Hour)
  }
  return eta
}
