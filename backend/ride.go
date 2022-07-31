package bike_app_be

import (
	"strconv"
	"strings"
)

type Ride struct {
	Departure              string `db:"departure" json:"departure"`
	Return                 string `db:"return" json:"return"`
	Departure_station_id   int    `db:"departure_station" json:"departure_station_id"`
	Return_station_id      int    `db:"return_station" json:"return_station_id"`
	Distance               int    `db:"distance" json:"distance"`
	Duration               int    `db:"duration" json:"duration"`
	Departure_station_name string `json:"departure_station_name"`
	Return_station_name    string `json:"return_station_name"`
}

func stringsToInts(strings []string) []int {
	ints := make([]int, 0, len(strings))
	for _, s := range strings {
		n, err := strconv.Atoi(s)
		if err != nil {
			ints = append(ints, -1)
		} else {
			ints = append(ints, n)
		}
	}
	return ints
}

func NewRidesFromDataSlice(station_ids map[int]struct{}, keys []string, data [][]string) [](*Ride) {
	info := make(map[string][]int)

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

	const min_duration = 10
	const min_distance = 10

	rideFilter := func(ride *Ride, i int) bool {
		if _, ok := station_ids[ride.Departure_station_id]; !ok {
			info["skipped"] = append(info["skipped"], i)
			info["station_id_err"] = append(info["station_id_err"], ride.Departure_station_id)
			return false
		}
		if _, ok := station_ids[ride.Return_station_id]; !ok {
			info["skipped"] = append(info["skipped"], i)
			info["station_id_err"] = append(info["station_id_err"], ride.Return_station_id)
			return false
		}
		if ride.Duration < min_duration {
			return false
		}
		if ride.Distance < min_distance {
			return false
		}
		return true
	}

	//fmt.Printf("%d,%d,%d,%d,%d,%d\n", d, r, did, rid, dis, dur)
	rides := make([](*Ride), 0, len(data))
	for i, row := range data {
		ints := stringsToInts(row)

		ride := &Ride{
			Departure:            row[d],
			Return:               row[r],
			Departure_station_id: ints[did],
			Return_station_id:    ints[rid],
			Distance:             ints[dis],
			Duration:             ints[dur],
		}
		if rideFilter(ride, i) {
			rides = append(rides, ride)
		}

	}
	//fmt.Println(info)
	return rides
}

type RideService interface {
	CreateRide(r *Ride) error
	GetCount() (int, error)
	GetRides(p map[string][]string) ([]*Ride, error)
}
