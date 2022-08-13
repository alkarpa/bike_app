package database

import (
	"testing"

	"alkarpa.fi/bike_app_be"
)

func TestStationService(t *testing.T) {
	db, err := openSQL(test_user, test_password, test_db)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	station_service := NewStationService(db)

	/*
		t.Run("GetAll", func(t *testing.T) { // TODO: Actual data test
			station_service.GetAll()
		})
	*/

	t.Run("Insert 2 stations", func(t *testing.T) {
		stations := make([]*bike_app_be.Station, 2)
		//fmt.Println("Before assignment")
		stations[0] = &bike_app_be.Station{}
		stations[1] = &bike_app_be.Station{}
		//fmt.Println("Before insert call")
		err := station_service.InsertStations(stations)
		if err != nil {
			t.Error(err)
		}
	})
}
