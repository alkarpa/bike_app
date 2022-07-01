package csvreader

import (
	"testing"
)

func TestFileDoesNotExist(t *testing.T) {
	_, err := ReadFromCSV("test_data/no_such_file")
	if err == nil {
		t.Fatalf("expected an error")
	}
}

func TestSimpleRideFile(t *testing.T) {
	dat, err := ReadFromCSV("test_data/simple_ride.csv")
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Run("data_line[0] Departure station id is 1", func(t *testing.T) {
		expected := "1"
		received := dat["Departure station id"][0]
		if expected != received {
			t.Fatalf("expected %s does not equal received %s", expected, received)
		}
	})
	t.Run("data_line[1] Return station id is 42", func(t *testing.T) {
		expected := "42"
		received := dat["Return station id"][1]
		if expected != received {
			t.Fatalf("expected %s does not equal received %s", expected, received)
		}
	})

}

func TestMissingDataFile(t *testing.T) {
	dat, err := ReadFromCSV("test_data/missing_data.csv")
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Run("data_line[2] Duration (sec.) is 62", func(t *testing.T) {
		expected := "62"
		received := dat["Duration (sec.)"][2]
		if expected != received {
			t.Fatalf("expected %s does not equal received %s", expected, received)
		}
	})
	t.Run("data_line[2] Duration (sec.) is 62", func(t *testing.T) {
		expected := "62"
		received := dat["Duration (sec.)"][2]
		if expected != received {
			t.Fatalf("expected %s does not equal received %s", expected, received)
		}
	})
}

func TestTooManyColumns(t *testing.T) {
	_, err := ReadFromCSV("test_data/too_many_columns.csv")
	if err != nil {
		t.Errorf(err.Error())
	}
}
