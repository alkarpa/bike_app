package csvreader

import (
	"testing"
)

func TestCsvReader(t *testing.T) {

	t.Run("File does not exist returns an error", func(t *testing.T) {
		_, err := ReadFromCSV("test_data/no_such_file")
		if err == nil {
			t.Fatalf("expected an error")
		}
	})

	t.Run("Simple ride file - data_line[0] Departure station id is 1", func(t *testing.T) {
		dat, err := ReadFromCSV("test_data/simple_ride.csv")
		if err != nil {
			t.Errorf(err.Error())
		}
		expected := "1"
		received := dat["Departure station id"][0]
		if expected != received {
			t.Fatalf("expected %s does not equal received %s", expected, received)
		}
	})
	t.Run("Simple ride file - data_line[1] Return station id is 42", func(t *testing.T) {
		dat, err := ReadFromCSV("test_data/simple_ride.csv")
		if err != nil {
			t.Errorf(err.Error())
		}
		expected := "42"
		received := dat["Return station id"][1]
		if expected != received {
			t.Fatalf("expected %s does not equal received %s", expected, received)
		}
	})

	t.Run("Missing data ride file - data_line[2] Duration (sec.) is 62", func(t *testing.T) {
		dat, err := ReadFromCSV("test_data/missing_data.csv")
		if err != nil {
			t.Errorf(err.Error())
		}
		expected := "62"
		received := dat["Duration (sec.)"][2]
		if expected != received {
			t.Fatalf("expected %s does not equal received %s", expected, received)
		}
	})

}
