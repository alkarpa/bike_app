package database

import "testing"

func TestDatabaseFailedOpen(t *testing.T) {
	_, err := openSQL("", "___", "___")
	if err == nil {
		t.Error(err)
	}
}
