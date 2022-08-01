package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"alkarpa.fi/bike_app_be"
)

func TestStation(t *testing.T) {
	ts := OpenTestServer(t)

	url := "/station"

	t.Run("GetAll", func(t *testing.T) {

		ts.StationService.GetAllFn = func() ([]*bike_app_be.Station, error) {
			all := make([]*bike_app_be.Station, 0, 3)
			for i := 0; i < 3; i++ {
				station := &bike_app_be.Station{Id: i}
				all = append(all, station)
			}
			return all, nil
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.getStations())

		handler.ServeHTTP(rr, req)

		t.Run("Status code", func(t *testing.T) {
			if expected, received := http.StatusOK, rr.Code; expected != received {
				t.Errorf("Expected %v, received %v", expected, received)
			}
		})
		t.Run("Returns 3", func(t *testing.T) {
			result := []bike_app_be.Station{}
			if err := json.NewDecoder(rr.Body).Decode(&result); err != nil {
				t.Error(err)
			}
			expected := 3
			received := len(result)
			if expected != received {
				t.Errorf("Expected %v, received %v", expected, received)
			}
		})

	})
	t.Run("GetAll DB error returns internal server error", func(t *testing.T) {

		ts.StationService.GetAllFn = func() ([]*bike_app_be.Station, error) {

			return nil, errors.New("test error")
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ts.getStations())

		handler.ServeHTTP(rr, req)

		expected := http.StatusInternalServerError
		received := rr.Code
		if expected != received {
			t.Errorf("Expected %v, received %v", expected, received)
		}

	})
	t.Run("GetStats", func(t *testing.T) {

		ts.StationService.GetDetailsFn = func(id int) (*bike_app_be.Stats, error) {
			if id == 1 {
				st := make(map[string]*bike_app_be.Stat_count_avg)
				st["1"] = &bike_app_be.Stat_count_avg{Count: 5}
				return &bike_app_be.Stats{Departing: st}, nil
			} else {
				return nil, errors.New("testing error")
			}
		}

		t.Run("good request", func(t *testing.T) {
			path := url + "/1"
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			ts.router.ServeHTTP(rr, req)

			if expected, received := http.StatusOK, rr.Code; expected != received {
				t.Fatalf("Expected %v, received %v", expected, received)
			}

			t.Run("Content", func(t *testing.T) {
				result := bike_app_be.Stats{}
				if err := json.NewDecoder(rr.Body).Decode(&result); err != nil {
					t.Error(err)
				}
				count := result.Departing["1"].Count
				expected := 5
				if count != expected {
					t.Errorf("Expected %v, received %v", expected, count)
				}
			})
		})
		t.Run("id not found", func(t *testing.T) {
			req, err := http.NewRequest("GET", url+"/2000", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			ts.router.ServeHTTP(rr, req)

			t.Run("Status code", func(t *testing.T) {
				if expected, received := http.StatusInternalServerError, rr.Code; expected != received {
					t.Errorf("Expected %v, received %v", expected, received)
				}
			})
		})
		t.Run("bad id", func(t *testing.T) {
			req, err := http.NewRequest("GET", url+"/test", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			ts.router.ServeHTTP(rr, req)

			t.Run("Status code", func(t *testing.T) {
				if expected, received := http.StatusBadRequest, rr.Code; expected != received {
					t.Errorf("Expected %v, received %v", expected, received)
				}
			})
		})
	})
}
