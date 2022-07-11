package database

import (
	"database/sql"
	"fmt"
	"strings"

	"alkarpa.fi/bike_app_be"
)

type RideService struct {
	db *sql.DB
}

func NewRideService(db *sql.DB) *RideService {
	return &RideService{db: db}
}

func statementValues(number_of_values uint) string {
	var qms = make([]string, 0, number_of_values)
	for i := 0; i < int(number_of_values); i++ {
		qms = append(qms, "?")
	}
	return fmt.Sprintf("(%s)", strings.Join(qms, ","))
}

func (rs *RideService) CreateRides(rides [](*bike_app_be.Ride)) error {
	// duration is temporarily used to count duplicates
	const sql_insert = "INSERT INTO ride  VALUES %s ON DUPLICATE KEY UPDATE duration=duration+1" //departure=departure"
	const number_of_values = 6
	stmt_values := statementValues(number_of_values)

	valueStrings := make([]string, 0, len(rides))
	values := make([]interface{}, 0, len(rides)*number_of_values)

	for _, ride := range rides {

		valueStrings = append(valueStrings, stmt_values)

		// duration is temporarily used to count duplicates
		values = append(values, ride.Departure, ride.Return, ride.Departure_station_id, ride.Return_station_id, ride.Distance, 0) // ride.Duration)
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
	return nil
}
