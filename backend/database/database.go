package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	csv "alkarpa.fi/csv_reader"
)

func openSQL(dbUser string, dbPassword string, dbName string) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%v:%v@/%v", dbUser, dbPassword, dbName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func OpenSQL() (*sql.DB, error) {
	dbUser, dbPassword, dbName := "bike_app_user", "bike_pw" /*os.Getenv("bike_app_PW")*/, "bike_app"
	return openSQL(dbUser, dbPassword, dbName)
}

// Initial placeholder for testing and base for future modifications
func ImportFromCSVs() error {
	fmt.Println("Test CSV")

	db, err := OpenSQL()
	if err != nil {
		return err
	}
	if err = csvStations(db); err != nil {
		return err
	}
	if err = csvRides(db); err != nil {
		return err
	}

	return nil
}

func csvStations(db *sql.DB) error {
	keys, data, err := csv.ReadFromCSV("../Helsingin_ja_Espoon_kaupunkipyöräasemat_avoin.csv") // TODO: remove hardcoded path
	if err != nil {
		return err
	}
	id := csv.GetKeyIndex("ID", keys)
	name := csv.GetKeyIndex("Nimi", keys)

	rows := make([][]string, 0)
	for _, row := range data {
		rows = append(rows, []string{row[id], row[name]})
	}
	station_service := NewStationService(db)
	return station_service.InsertStations(rows)
}
func csvRides(db *sql.DB) error {
	// TODO
	_, _, err := csv.ReadFromCSV("../2021-05.csv")
	if err != nil {
		return err
	}

	return nil
}
