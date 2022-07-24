package database

import (
	"database/sql"
	"fmt"
	"strings"

	"alkarpa.fi/bike_app_be"
)

type StationService struct {
	db *sql.DB
}

func NewStationService(db *sql.DB) *StationService {
	return &StationService{db: db}
}

func (s *StationService) GetAll() ([]*bike_app_be.Station, error) {
	rows, err := s.db.Query("SELECT * FROM station")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lang_fields, err := s.GetAllStationLangFields()
	if err != nil {
		return nil, err
	}
	stations := []*bike_app_be.Station{}

	for rows.Next() {
		var station = &bike_app_be.Station{}
		err = rows.Scan(&station.Id,
			&station.Operator,
			&station.Capacity,
			&station.X,
			&station.Y)
		if err != nil {
			return nil, err
		}
		station.Text = bike_app_be.BuildTextFromLangFields(lang_fields[station.Id])
		stations = append(stations, station)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return stations, nil
}

func (s *StationService) getStatCountAVG(statmap *map[string]*bike_app_be.Stat_count_avg, query string) error {
	rows, err := s.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		yearmonth := ""
		r := &bike_app_be.Stat_count_avg{}
		if err = rows.Scan(&yearmonth, &r.Count, &r.Average_distance); err != nil {
			return err
		}
		(*statmap)[yearmonth] = r
	}
	if err = rows.Err(); err != nil {
		return err
	}
	return nil
}
func (s *StationService) getTopConnections(statmap *map[string]*bike_app_be.Stat_count_avg, query string) error {
	rows, err := s.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		yearmonth := ""
		r := bike_app_be.Stat_connection{}
		if err = rows.Scan(&yearmonth, &r.Station_id, &r.Count); err != nil {
			return err
		}
		if _, ok := (*statmap)[yearmonth]; ok {
			(*statmap)[yearmonth].Top_connections = append((*statmap)[yearmonth].Top_connections, r)
		}
	}
	if err = rows.Err(); err != nil {
		return err
	}
	return nil
}

func (s *StationService) GetDetails(id int) (*bike_app_be.Stats, error) {
	departing := make(map[string]*bike_app_be.Stat_count_avg)
	returning := make(map[string]*bike_app_be.Stat_count_avg)

	const dep, ret = "departure", "return"
	helpers := []*stat_query_helper{
		{target: &departing, a: dep, b: ret, id: id},
		{target: &returning, a: ret, b: dep, id: id},
	}

	queries := []stat_query{}
	for _, h := range helpers {
		queries = append(queries,
			&query_count_average{h},
			&query_top_connections{h},
		)
	}

	for _, query := range queries {
		for _, monthly := range []bool{false, true} {
			if err := query.make_query(s, monthly); err != nil {
				return nil, err
			}
		}

	}

	stats := &bike_app_be.Stats{
		Departing: departing,
		Returning: returning,
	}
	return stats, nil

}

func (s *StationService) GetAllStationLangFields() (map[int][]bike_app_be.StationLangField, error) {
	rows, err := s.db.Query("SELECT * FROM station_lang_field")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fields := make(map[int][]bike_app_be.StationLangField)

	for rows.Next() {
		var id = 0
		var field = bike_app_be.StationLangField{}
		err = rows.Scan(&id, &field.Lang, &field.Key, &field.Value)
		if err != nil {
			return nil, err
		}
		fields[id] = append(fields[id], field)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return fields, nil
}

func (s *StationService) InsertStations(stations []*bike_app_be.Station) error {
	const sql_insert = "INSERT INTO station VALUES %s ON DUPLICATE KEY UPDATE id=id"
	const number_of_values = 5
	stmt_values := statementValues(number_of_values)

	valueStrings := make([]string, 0, len(stations))
	values := make([]interface{}, 0, len(stations)*number_of_values)

	for _, station := range stations {
		valueStrings = append(valueStrings, stmt_values)
		values = append(values, station.Id, station.Operator, station.Capacity, station.X, station.Y)
	}
	insert_query := fmt.Sprintf(sql_insert, strings.Join(valueStrings, ","))

	stmt, err := s.db.Prepare(insert_query)
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

func (s *StationService) insertStationLangFields(stations []*bike_app_be.Station) error {
	const sql_insert = "INSERT INTO station_lang_field VALUES %s ON DUPLICATE KEY UPDATE id=id"
	const number_of_values = 4
	stmt_values := statementValues(number_of_values)

	valueStrings := make([]string, 0, len(stations))
	values := make([]interface{}, 0, len(stations)*number_of_values)

	for _, station := range stations {
		for _, field := range station.LangFields {
			valueStrings = append(valueStrings, stmt_values)
			values = append(values, station.Id, field.Lang, field.Key, field.Value)
		}
	}

	insert_query := fmt.Sprintf(sql_insert, strings.Join(valueStrings, ","))

	stmt, err := s.db.Prepare(insert_query)
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
