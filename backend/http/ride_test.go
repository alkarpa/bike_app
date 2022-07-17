package http

import (
	"net/http"
	"testing"

	"alkarpa.fi/bike_app_be"
)

func TestRide(t *testing.T) {
	ts := OpenTestServer(t)
	defer ts.Close()

	t.Run("GetRides", func(t *testing.T) {

		ts.RideService.GetRidesFn = func() ([]*bike_app_be.Ride, error) {
			return nil, nil
		}

		req, err := http.NewRequest("GET", "http://localhost:8080/ride/", nil)
		if err != nil {
			t.Fatal(err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer res.Body.Close()
		t.Run("Status code", func(t *testing.T) {
			if received, expected := res.StatusCode, http.StatusOK; received != expected {
				t.Fatalf("Expected %v, received %v", expected, received)
			}
		})
		t.Run("Content", func(t *testing.T) {

		})

	})
}
