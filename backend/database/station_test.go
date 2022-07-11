package database

import (
	"testing"
)

func TestStationService(t *testing.T) {
	db, err := openSQL("bike_app_test_user", "bike_pw", "bike_app_test")
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
		stations := make([][]string, 2)
		//fmt.Println("Before assignment")
		stations[0] = []string{"1", "Boston"}
		stations[1] = []string{"2", "Nyark"}
		//fmt.Println("Before insert call")
		err := station_service.InsertStations(stations)
		if err != nil {
			t.Error(err)
		}
	})
}
