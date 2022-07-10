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

	t.Run("GetAll", func(t *testing.T) { // TODO: Actual data test
		station_service.GetAll()
	})
}
