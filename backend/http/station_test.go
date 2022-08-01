package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"alkarpa.fi/bike_app_be"
)

func TestStation(t *testing.T) {
	ts := OpenTestServer(t)
	defer ts.Close()

	url := test_url + "/station/"

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
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer res.Body.Close()

		t.Run("Status code", func(t *testing.T) {
			if expected, received := http.StatusOK, res.StatusCode; expected != received {
				t.Errorf("Expected %v, received %v", expected, received)
			}
		})
		t.Run("Returns 3", func(t *testing.T) {
			result := []bike_app_be.Station{}
			if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
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
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		defer res.Body.Close()
		expected := http.StatusInternalServerError
		received := res.StatusCode
		if expected != received {
			t.Errorf("Expected %v, received %v", expected, received)
		}

	})
	t.Run("GetStats", func(t *testing.T) {

		ts.StationService.GetDetailsFn = func(id int) (*bike_app_be.Stats, error) {
			if id == 1 {
				return &bike_app_be.Stats{}, nil
			} else {
				return nil, errors.New("testing error")
			}
		}

		t.Run("good request", func(t *testing.T) {
			req, err := http.NewRequest("GET", url+"1", nil)
			if err != nil {
				t.Fatal(err)
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			t.Run("Status code", func(t *testing.T) {
				if expected, received := http.StatusOK, res.StatusCode; expected != received {
					t.Errorf("Expected %v, received %v", expected, received)
				}
			})
		})
		t.Run("bad request", func(t *testing.T) {
			req, err := http.NewRequest("GET", url+"2000", nil)
			if err != nil {
				t.Fatal(err)
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			t.Run("Status code", func(t *testing.T) {
				if expected, received := http.StatusInternalServerError, res.StatusCode; expected != received {
					t.Errorf("Expected %v, received %v", expected, received)
				}
			})
		})
		/*
			t.Run("Returns 3", func(t *testing.T) {
				result := []bike_app_be.Station{}
				if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
					t.Error(err)
				}
				expected := 3
				received := len(result)
				if expected != received {
					t.Errorf("Expected %v, received %v", expected, received)
				}
			})*/
	})
}
