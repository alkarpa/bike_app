package database

import (
	"database/sql"
	"fmt"
	"strings"

	"alkarpa.fi/bike_app_be"
)

const unset_ride_count = -1

type RideService struct {
	db         *sql.DB
	ride_count int
}

func NewRideService(db *sql.DB) *RideService {
	return &RideService{
		db:         db,
		ride_count: unset_ride_count,
	}
}

func statementValues(number_of_values uint) string {
	var qms = make([]string, 0, number_of_values)
	for i := 0; i < int(number_of_values); i++ {
		qms = append(qms, "?")
	}
	return fmt.Sprintf("(%s)", strings.Join(qms, ","))
}

func (rs *RideService) GetCount() (int, error) {
	// optimization trick
	if rs.ride_count != unset_ride_count {
		return rs.ride_count, nil
	}

	rows, err := rs.db.Query("SELECT COUNT(*) FROM ride")
	if err != nil {
		return unset_ride_count, err
	}
	defer rows.Close()
	var count int
	if rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return unset_ride_count, err
		}
	}
	if err = rows.Err(); err != nil {
		return unset_ride_count, err
	}
	return count, nil
}

// Call after operations that change the total count of rides.
func (rs *RideService) countHasChanged() {
	rs.ride_count = unset_ride_count
}

func (rs *RideService) CreateRides(rides [](*bike_app_be.Ride)) error {
	const sql_insert = "INSERT INTO ride  VALUES %s ON DUPLICATE KEY UPDATE departure=departure"
	const number_of_values = 6
	stmt_values := statementValues(number_of_values)

	valueStrings := make([]string, 0, len(rides))
	values := make([]interface{}, 0, len(rides)*number_of_values)

	for _, ride := range rides {

		valueStrings = append(valueStrings, stmt_values)

		values = append(values, ride.Departure, ride.Return, ride.Departure_station_id, ride.Return_station_id, ride.Distance, ride.Duration)
	}
	insert_query := fmt.Sprintf(sql_insert, strings.Join(valueStrings, ","))

	stmt, err := rs.db.Prepare(insert_query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(values...)
	if err != nil {
		return err
	}
	rs.countHasChanged()
	return nil
}

func (rs *RideService) CreateRide(ride *bike_app_be.Ride) error {
	return nil
}

func (rs *RideService) GetRides(parameters map[string][]string) ([]*bike_app_be.Ride, error) {

	sqf := newRideSelectQueryFriend(parameters)

	//fmt.Println(sqf.buildQuery())

	rows, err := rs.db.Query(sqf.buildQuery(), sqf.values...)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rides := []*bike_app_be.Ride{}

	for rows.Next() {
		var ride = &bike_app_be.Ride{}
		err := rows.Scan(&ride.Departure, &ride.Return, &ride.Departure_station_id, &ride.Return_station_id,
			&ride.Distance, &ride.Duration, &ride.Departure_station_name, &ride.Return_station_name)
		if err != nil {
			return nil, err
		}
		rides = append(rides, ride)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rides, nil
}
