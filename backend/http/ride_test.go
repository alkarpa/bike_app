package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"alkarpa.fi/bike_app_be"
)

func TestRide(t *testing.T) {
	ts := OpenTestServer(t)
	defer ts.Close()

	url := test_url + "/ride/"

	t.Run("GetRides", func(t *testing.T) {

		ts.RideService.GetRidesFn = func() ([]*bike_app_be.Ride, error) {
			rides := make([]*bike_app_be.Ride, 0, 2)
			for i := 0; i < 2; i++ {
				ride := &bike_app_be.Ride{}
				rides = append(rides, ride)
			}
			return rides, nil
		}

		req, err := http.NewRequest("GET", url, nil)
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
		t.Run("Returns len 2 array", func(t *testing.T) {
			result := []bike_app_be.Ride{}
			if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
				t.Error(err)
			}
			expected := 2
			received := len(result)
			if expected != received {
				t.Errorf("Expected %v, received %v", expected, received)
			}
		})

	})
	t.Run("GetRides db error", func(t *testing.T) {

		ts.RideService.GetRidesFn = func() ([]*bike_app_be.Ride, error) {
			return nil, errors.New("test error")
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatal(err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer res.Body.Close()
		t.Run("Status code", func(t *testing.T) {
			if received, expected := res.StatusCode, http.StatusInternalServerError; received != expected {
				t.Fatalf("Expected %v, received %v", expected, received)
			}
		})
	})
}
