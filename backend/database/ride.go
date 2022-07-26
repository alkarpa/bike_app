package database

import (
	"database/sql"
	"fmt"
	"strconv"
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
	return nil
}

func (rs *RideService) CreateRide(ride *bike_app_be.Ride) error {
	return nil
}

func (rs *RideService) getParamsPage(parameters map[string][]string) string {
	const page_size = 100
	default_limit := fmt.Sprintf("LIMIT %d", page_size)
	if page_params, found := parameters["page"]; found {
		page, err := strconv.Atoi(page_params[0])
		if err != nil || page < 1 {
			return default_limit
		}
		const page_limit_offset = "LIMIT %d OFFSET %d"
		return fmt.Sprintf(page_limit_offset, page_size, page_size*page)
	}
	return default_limit
}

func (rs *RideService) getParamsOrdering(parameters map[string][]string) string {
	if order_params, found := parameters["order"]; found {
		accepted_values := []string{
			"departure",
			"return",
			"departure_station",
			"return_station",
			"distance",
			"duration",
		}
		order_string := "ORDER BY %s"
		orderings := make([]string, 0, len(order_params))
		for _, ordering := range order_params {
			parts := strings.Split(ordering, "_")
			for _, accepted_value := range accepted_values {
				if parts[0] == accepted_value {
					query_ord := accepted_value
					if len(parts) > 1 && parts[1] == "desc" {
						query_ord += " desc"
					}
					orderings = append(orderings, query_ord)
					break
				}
			}
		}
		if len(orderings) > 0 {
			return fmt.Sprintf(order_string, strings.Join(orderings, ","))
		}
	}
	return ""
}

func (rs *RideService) GetRides(parameters map[string][]string) ([]*bike_app_be.Ride, error) {
	//fmt.Println(rs.getParamsOrdering(parameters))
	//fmt.Println(rs.getParamsPage(parameters))

	// TODO: filtering
	const select_values = "departure, `return`, departure_station, return_station, distance, duration"

	query := fmt.Sprintf("SELECT %[1]s FROM ride ", select_values)

	values := make([]interface{}, 0)
	if search, search_found := parameters["search"]; search_found {
		query += " INNER JOIN " +
			"(SELECT DISTINCT id FROM station_lang_field WHERE station_lang_field.lang = 'fi' " +
			"AND station_lang_field.value LIKE CONCAT('%',?,'%') ) a " +
			"ON a.id IN (departure_station, return_station) "
		values = append(values, search[0])
		fmt.Println(values)
	}
	query += rs.getParamsOrdering(parameters) + " "
	query += rs.getParamsPage(parameters) + " "
	rows, err := rs.db.Query(query, values...)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rides := []*bike_app_be.Ride{}

	for rows.Next() {
		var ride = &bike_app_be.Ride{}

		err := rows.Scan(&ride.Departure, &ride.Return, &ride.Departure_station_id, &ride.Return_station_id, &ride.Distance, &ride.Duration)
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
