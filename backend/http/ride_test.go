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

	t.Run("GetCount ok", func(t *testing.T) {
		ts.RideService.GetCountFn = func() (int, error) {
			return 2, nil
		}

		t.Run("GetRides", func(t *testing.T) {

			ts.RideService.GetRidesFn = func(p map[string][]string) ([]*bike_app_be.Ride, error) {
				rides := make([]*bike_app_be.Ride, 0, 2)
				for i := 0; i < 2; i++ {
					rides = append(rides, &bike_app_be.Ride{
						Departure_station_id: i,
						Return_station_id:    i + 2,
					})
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
			t.Run("Read body", func(t *testing.T) {
				result := &struct {
					Count int                 `json:"count"`
					Rides []*bike_app_be.Ride `json:"rides"`
				}{}
				if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
					t.Error(err)
				}
				t.Run("len(data) is 2", func(t *testing.T) {
					expected := 2
					received := len(result.Rides)
					if expected != received {
						t.Errorf("Expected %v, received %v", expected, received)
					}
				})
				t.Run("count is 2", func(t *testing.T) {
					expected := 2
					received := result.Count
					if expected != received {
						t.Errorf("Expected %v, received %v", expected, received)
					}
				})
				t.Run("Departure stations are 0 and 1", func(t *testing.T) {
					expected := []int{0, 1}
					received := result.Rides
					for i, _ := range expected {
						if expected[i] != received[i].Departure_station_id {
							t.Errorf("Expected %v, received %v", expected, received)
						}

					}
				})
			})

		})
		t.Run("GetRides db error", func(t *testing.T) {

			ts.RideService.GetRidesFn = func(p map[string][]string) ([]*bike_app_be.Ride, error) {
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

	})
	t.Run("GetCount bad", func(t *testing.T) {
		ts.RideService.GetCountFn = func() (int, error) {
			return -1, errors.New("testing error")
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
