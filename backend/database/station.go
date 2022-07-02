package database

import "database/sql"

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
