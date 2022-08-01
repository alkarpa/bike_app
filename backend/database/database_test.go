package database

import (
	"database/sql"
	"io/ioutil"
	"strings"
	"testing"
)

const test_db = "bike_app_test"
const test_user = "bike_app_test_user"
const test_password = "bike_pw"

func TestDatabaseFailedOpen(t *testing.T) {
	_, err := openSQL("", "___", "___")
	if err == nil {
		t.Error(err)
	}
}

func initializeDatabaseForTests(db *sql.DB) error {
	// Use the same sql file as is used for the actual database
	query_bits, err := ioutil.ReadFile("../../bike_app.sql")
	if err != nil {
		return err
	}
	query := string(query_bits)
	queries := strings.Split(query, ";")
	for _, q := range queries {
		trim_q := strings.TrimSpace(q)
		if len(trim_q) > 0 {
			_, err = db.Exec(trim_q)
			if err != nil {
				return err
			}
		}

	}

	importer := CSVImporter{
		Path:     "../../test_data/",
		Verbose:  false,
		Database: db,
	}
	if err = importer.ImportFromCSVs(); err != nil {
		return err
	}

	return nil // errors.New("just for test print")
}

func TestDatabase(t *testing.T) {
	db, err := openSQL(test_user, test_password, test_db)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	if err := initializeDatabaseForTests(db); err != nil {
		t.Error(err)
	}

	t.Run("Rides", func(t *testing.T) {
		ride_service := NewRideService(db)

		t.Run("GetCount", func(t *testing.T) {
			count, err := ride_service.GetCount()
			if err != nil {
				t.Error(err)
			}
			received := count
			expected := 3
			if received != expected {
				t.Errorf("expected '%d', got '%d'", expected, received)
			}
		})

		t.Run("GetRides", func(t *testing.T) {
			rides, err := ride_service.GetRides(map[string][]string{})
			if err != nil {
				t.Error(err)
			}
			t.Run("len(rides)=3", func(t *testing.T) {
				received := len(rides)
				expected := 3
				if received != expected {
					t.Errorf("expected '%d', got '%d'", expected, received)
				}
			})
			t.Run("rides has one ride from 1 to 2", func(t *testing.T) {
				count := 0
				for _, ride := range rides {
					if ride.Departure_station_id == 1 && ride.Return_station_id == 2 {
						count++
					}
				}
				expected := 1
				if count != expected {
					t.Errorf("expected '%d', got '%d'", expected, count)
				}
			})
			t.Run("rides has one ride from 2 to 1", func(t *testing.T) {
				count := 0
				for _, ride := range rides {
					if ride.Departure_station_id == 2 && ride.Return_station_id == 1 {
						count++
					}
				}
				expected := 1
				if count != expected {
					t.Errorf("expected '%d', got '%d'", expected, count)
				}
			})
		})
		t.Run("GetRides search Etsi", func(t *testing.T) {
			rides, err := ride_service.GetRides(map[string][]string{"search": {"Etsi"}})
			if err != nil {
				t.Error(err)
			}
			t.Run("len(rides)=1", func(t *testing.T) {
				received := len(rides)
				expected := 1
				if received != expected {
					t.Errorf("expected '%d', got '%d'", expected, received)
				}
			})
			t.Run("rides has one ride from 3 to 3", func(t *testing.T) {
				count := 0
				for _, ride := range rides {
					if ride.Departure_station_id == 3 && ride.Return_station_id == 3 {
						count++
					}
				}
				expected := 1
				if count != expected {
					t.Errorf("expected '%d', got '%d'", expected, count)
				}
			})
			t.Run("ride has distance of 2000", func(t *testing.T) {
				received := rides[0].Distance
				expected := 2000
				if received != expected {
					t.Errorf("expected '%d', got '%d'", expected, received)
				}
			})
			t.Run("ride has duration of 60", func(t *testing.T) {
				received := rides[0].Duration
				expected := 60
				if received != expected {
					t.Errorf("expected '%d', got '%d'", expected, received)
				}
			})

		})
		t.Run("GetRides order duration_desc", func(t *testing.T) {
			rides, err := ride_service.GetRides(map[string][]string{"order": {"duration_desc"}})
			if err != nil {
				t.Error(err)
			}
			t.Run("len(rides)=3", func(t *testing.T) {
				received := len(rides)
				expected := 3
				if received != expected {
					t.Errorf("expected '%d', got '%d'", expected, received)
				}
			})
			t.Run("rides durations in descending order", func(t *testing.T) {
				last_duration := rides[0].Duration
				for i, ride := range rides {
					if last_duration < ride.Duration {
						t.Errorf("expected descending order, %d:%d is less than %d:%d", i-1, last_duration, i, ride.Duration)
					}
					last_duration = ride.Duration
				}
			})
			t.Run("largest duration is 72", func(t *testing.T) {
				received := rides[0].Duration
				expected := 72
				if received != expected {
					t.Errorf("expected '%d', got '%d'", expected, received)
				}
			})

		})

	})

	t.Run("Stations", func(t *testing.T) {
		station_service := NewStationService(db)

		t.Run("GetAll", func(t *testing.T) {
			stations, err := station_service.GetAll()
			if err != nil {
				t.Error(err)
			}
			t.Run("len(stations)=3", func(t *testing.T) {
				received := len(stations)
				expected := 3
				if received != expected {
					t.Errorf("expected '%d', got '%d'", expected, received)
				}
			})

		})

		t.Run("GetDetails", func(t *testing.T) {
			stats, err := station_service.GetDetails(1)
			if err != nil {
				t.Error(err)
			}
			t.Run("Departing", func(t *testing.T) {
				departing := stats.Departing
				t.Run("7 has 1 departing", func(t *testing.T) {
					received := departing["2022-7"].Count
					expected := 1
					if received != expected {
						t.Errorf("expected '%d', got '%d'", expected, received)
					}
				})
				t.Run("8 is empty", func(t *testing.T) {
					received := departing["2022-8"]
					if received != nil {
						t.Errorf("expected nil, got '%v'", received)
					}
				})
				t.Run("all has 1 departing", func(t *testing.T) {
					received := departing["all"].Count
					expected := 1
					if received != expected {
						t.Errorf("expected '%d', got '%d'", expected, received)
					}
				})
			})
			t.Run("Returning", func(t *testing.T) {
				returning := stats.Returning
				t.Run("8 has 1 returning", func(t *testing.T) {
					received := returning["2022-8"].Count
					expected := 1
					if received != expected {
						t.Errorf("expected '%d', got '%d'", expected, received)
					}
				})
				t.Run("7 is empty", func(t *testing.T) {
					received := returning["2022-7"]
					if received != nil {
						t.Errorf("expected nil, got '%v'", received)
					}
				})
				t.Run("all has 1 returning", func(t *testing.T) {
					received := returning["all"].Count
					expected := 1
					if received != expected {
						t.Errorf("expected '%d', got '%d'", expected, received)
					}
				})
			})
		})

	})

}
