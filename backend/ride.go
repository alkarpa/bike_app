package bike_app_be

import (
	"strconv"
	"strings"
)

type Ride struct {
	Departure            string `db:"departure" json:"departure"`
	Return               string `db:"return" json:"return"`
	Departure_station_id int    `db:"departure_station" json:"departure_station_id"`
	Return_station_id    int    `db:"return_station" json:"return_station_id"`
	Distance             int    `db:"distance" json:"distance"`
	Duration             int    `db:"duration" json:"duration"`
}

func stringsToInts(strings []string) []int {
	ints := make([]int, 0, len(strings))
	for _, s := range strings {
		n, err := strconv.Atoi(s)
		if err != nil {
			ints = append(ints, -1)
		}
		ints = append(ints, n)
	}
	return ints
}

func NewRidesFromDataSlice(keys []string, data [][]string) [](*Ride) {
	//fmt.Printf("NewRidesFromDataSlice, len(data):%d\n", len(data))
	d, r, did, rid, dis, dur := -1, -1, -1, -1, -1, -1
	for i, key := range keys {
		switch strings.TrimSpace(key) {
		case "Departure":
			d = i
		case "Return":
			r = i
		case "Departure station id":
			did = i
		case "Return station id":
			rid = i
		case "Covered distance (m)":
			dis = i
		case "Duration (sec.)":
			dur = i
		}
	}
	//fmt.Printf("%d,%d,%d,%d,%d,%d\n", d, r, did, rid, dis, dur)
	rides := make([](*Ride), 0, len(data))
	for _, row := range data {
		ints := stringsToInts(row)

		rides = append(rides, &Ride{
			Departure:            row[d],
			Return:               row[r],
			Departure_station_id: ints[did],
			Return_station_id:    ints[rid],
			Distance:             ints[dis],
			Duration:             ints[dur],
		})
	}
	return rides
}

type RideService interface {
	CreateRide(r *Ride) error
	GetRides() ([]*Ride, error)
}
