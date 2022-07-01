package csvreader

import (
	"testing"
)

func TestCsvReader(t *testing.T) {

	t.Run("File does not exist", func(t *testing.T) {
		_, err := ReadFromCSV("test_data/no_such_file")
		if err == nil {
			t.Fatalf("expected an error")
		}
	})

	t.Run("Simple ride file", func(t *testing.T) {
		received, err := ReadFromCSV("test_data/simple_ride.csv")
		if err != nil {
			t.Errorf(err.Error())
		}
		expected := "1"
		if expected != received["Departure station id"][0] {
			t.Fatalf("Expected Id does not equal received Id")
		}
	})

}
