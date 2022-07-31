package mock

import "alkarpa.fi/bike_app_be"

type RideService struct {
	CreateRideFn func(r *bike_app_be.Ride) error
	GetCountFn   func() (int, error)
	GetRidesFn   func(map[string][]string) ([]*bike_app_be.Ride, error)
}

func (rs *RideService) CreateRide(r *bike_app_be.Ride) error {
	return rs.CreateRideFn(r)
}

func (rs *RideService) GetRides(p map[string][]string) ([]*bike_app_be.Ride, error) {
	return rs.GetRidesFn(p)
}

func (rs *RideService) GetCount() (int, error) {
	return rs.GetCountFn()
}
