package csvreader

import (
	"testing"
)

func TestFileDoesNotExist(t *testing.T) {
	_, _, err := ReadFromCSV("test_data/no_such_file")
	if err == nil {
		t.Fatalf("expected an error")
	}
}

func TestSimpleRideFile(t *testing.T) {
	keys, dat, err := ReadFromCSV("test_data/simple_ride.csv")
	if err != nil {
		t.Errorf(err.Error())
	}
	depStatKey := GetKeyIndex("Departure station id", keys)
	retStatKey := GetKeyIndex("Return station id", keys)

	t.Run("dat has a length of 2", func(t *testing.T) {
		expected := 2
		received := len(dat)
		if expected != received {
			t.Fatalf("expected %d does not equal received %d", expected, received)
		}
	})
	t.Run("data_line[0] Departure station id is 1", func(t *testing.T) {
		expected := "1"
		received := dat[0][depStatKey]
		if expected != received {
			t.Fatalf("expected %s does not equal received %s", expected, received)
		}
	})
	t.Run("data_line[1] Return station id is 42", func(t *testing.T) {
		expected := "42"
		received := dat[1][retStatKey]
		if expected != received {
			t.Fatalf("expected %s does not equal received %s", expected, received)
		}
	})

}

func TestMissingDataFile(t *testing.T) {
	keys, dat, err := ReadFromCSV("test_data/missing_data.csv")
	if err != nil {
		t.Errorf(err.Error())
	}

	durKey := GetKeyIndex("Duration (sec.)", keys)

	t.Run("data_line[2] Duration (sec.) is 62", func(t *testing.T) {
		expected := "62"
		received := dat[2][durKey]
		if expected != received {
			t.Fatalf("expected %s does not equal received %s", expected, received)
		}
	})
	t.Run("data_line[2] Duration (sec.) is 62", func(t *testing.T) {
		expected := "62"
		received := dat[2][durKey]
		if expected != received {
			t.Fatalf("expected %s does not equal received %s", expected, received)
		}
	})
}

func TestTooManyColumns(t *testing.T) {
	_, _, err := ReadFromCSV("test_data/too_many_columns.csv")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestKeyIndex(t *testing.T) {
	keys, _, err := ReadFromCSV("test_data/simple_ride.csv")
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Run("not in file", func(t *testing.T) {
		received := GetKeyIndex("not in file", keys)
		expected := -1
		if received != expected {
			t.Fatalf("expected %d, received %d", expected, received)
		}
	})
	t.Run("Departure", func(t *testing.T) {
		received := GetKeyIndex("Departure", keys)
		expected := 0
		if received != expected {
			t.Fatalf("expected %d, received %d", expected, received)
		}
	})
	t.Run("Departure station id", func(t *testing.T) {
		received := GetKeyIndex("Departure station id", keys)
		expected := 2
		if received != expected {
			t.Fatalf("expected %d, received %d", expected, received)
		}
	})
	t.Run("Return station id", func(t *testing.T) {
		received := GetKeyIndex("Return station id", keys)
		expected := 3
		if received != expected {
			t.Fatalf("expected %d, received %d", expected, received)
		}
	})

}
