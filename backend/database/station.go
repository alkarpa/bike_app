package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type StationService struct {
	db *sql.DB
}

func NewStationService(db *sql.DB) *StationService {
	return &StationService{db: db}
}

func (s *StationService) GetAll() ([]string, error) {
	rows, err := s.db.Query("SELECT * FROM station")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		// TODO
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *StationService) InsertStations(rows [][]string) error {
	sqlInsert := "INSERT INTO station(id, name) VALUES "
	values := []interface{}{}
	const rowSQL = "(%d, \"%s\")"
	inserts := make([]string, len(rows))
	//fmt.Println("Before loop")
	for i, row := range rows {
		id, _ := strconv.Atoi(row[0])
		inserts[i] = fmt.Sprintf(rowSQL, id, row[1])
	}
	sqlInsert = sqlInsert + strings.Join(inserts, ",")
	sqlInsert = sqlInsert + " ON DUPLICATE KEY UPDATE id=id"
	//fmt.Println(sqlInsert)

	stmt, err := s.db.Prepare(sqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		return err
	}
	//fmt.Println(ins)
	return nil
}
