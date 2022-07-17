package mock

import "alkarpa.fi/bike_app_be"

type RideService struct {
	CreateRideFn func(r *bike_app_be.Ride) error
	GetRidesFn   func() ([]*bike_app_be.Ride, error)
}

func (rs *RideService) CreateRide(r *bike_app_be.Ride) error {
	return rs.CreateRideFn(r)
}

func (rs *RideService) GetRides() ([]*bike_app_be.Ride, error) {
	return rs.GetRidesFn()
}
