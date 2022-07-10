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

func ImportFromCSVs() error {
	fmt.Println("Test CSV")
	a, err := csv.ReadFromCSV("../Helsingin_ja_Espoon_kaupunkipyöräasemat_avoin.csv")
	if err != nil {
		return err
	}
	fmt.Println(len(a))
	return nil
}
