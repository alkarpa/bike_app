package mock

import "alkarpa.fi/bike_app_be"

type StationService struct {
	StationService bike_app_be.StationService

	GetAllFn func() ([]*bike_app_be.Station, error)
}

func (ss *StationService) GetAll() ([]*bike_app_be.Station, error) {
	return ss.GetAllFn()
}
