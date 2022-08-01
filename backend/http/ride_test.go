package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"alkarpa.fi/bike_app_be"
)

func TestRide(t *testing.T) {
	ts := OpenTestServer(t)
	path := "/ride"

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

			req, err := http.NewRequest(http.MethodGet, path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(ts.getRides())

			handler.ServeHTTP(rr, req)

			t.Run("Status code", func(t *testing.T) {
				if received, expected := rr.Code, http.StatusOK; received != expected {
					t.Fatalf("Expected %v, received %v", expected, received)
				}
			})
			t.Run("Read body", func(t *testing.T) {
				result := &struct {
					Count int                 `json:"count"`
					Rides []*bike_app_be.Ride `json:"rides"`
				}{}
				if err := json.NewDecoder(rr.Body).Decode(&result); err != nil {
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
					for i := range expected {
						if expected[i] != received[i].Departure_station_id {
							t.Errorf("Expected %v, received %v", expected, received[i].Departure_station_id)
						}

					}
				})
			})

		})
		t.Run("GetRides db error", func(t *testing.T) {

			ts.RideService.GetRidesFn = func(p map[string][]string) ([]*bike_app_be.Ride, error) {
				return nil, errors.New("test error")
			}

			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(ts.getRides())

			handler.ServeHTTP(rr, req)

			t.Run("Status code", func(t *testing.T) {
				if received, expected := rr.Code, http.StatusInternalServerError; received != expected {
					t.Fatalf("Expected %v, received %v", expected, received)
				}
			})
		})

	})
	t.Run("GetCount bad", func(t *testing.T) {
		ts.RideService.GetCountFn = func() (int, error) {
			return -1, errors.New("testing error")
		}
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.getRides())

		handler.ServeHTTP(rr, req)

		t.Run("Status code", func(t *testing.T) {
			if received, expected := rr.Code, http.StatusInternalServerError; received != expected {
				t.Fatalf("Expected %v, received %v", expected, received)
			}
		})
	})
}
