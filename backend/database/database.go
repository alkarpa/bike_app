package database

import (
	"database/sql"
	"fmt"
	"math"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"

	"alkarpa.fi/bike_app_be"
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
	fmt.Println("Import from CSV")

	db, err := OpenSQL()
	if err != nil {
		return err
	}

	fmt.Println("\n-=Stations")
	if err = csvStations(db); err != nil {
		return err
	}
	fmt.Println("\n-=Rides")
	if err = csvRides(db); err != nil {
		return err
	}

	return nil
}

func csvStations(db *sql.DB) error {
	const directory = "../data/station/"
	files, err := os.ReadDir(directory)
	if err != nil {
		return err
	}
	for _, f := range files {
		path := fmt.Sprintf("%s%s", directory, f.Name())
		fmt.Printf("Importing %s\n", path)
		if err := addStationsFromCsv(db, path); err != nil {
			fmt.Println(err.Error())
		}
	}
	return nil
}
func addStationsFromCsv(db *sql.DB, path string) error {
	keys, data, err := csv.ReadFromCSV(path)
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
	const directory = "../data/ride/"
	files, err := os.ReadDir(directory)
	if err != nil {
		return err
	}
	ride_file_names := make([]string, 0, len(files))
	for _, f := range files {
		ride_file_names = append(ride_file_names, fmt.Sprintf("%s%s", directory, f.Name()))
	}
	//fmt.Println("")
	ride_service := NewRideService(db)
	for _, path := range ride_file_names {
		import_info := fmt.Sprintf("Importing '%s'", path)
		fmt.Print(import_info)
		keys, data, err := csv.ReadFromCSV(path)
		if err != nil {
			return err
		}

		var wg sync.WaitGroup

		const batch_size = 2048
		const max_consecutive = 24
		consecutives_channel := make(chan int, max_consecutive)
		active_routines := 0

		print_update := func(i int) {
			fmt.Printf("\r%s - [%7d / %7d]", import_info, i, len(data))
		}

		for i := 0; i < len(data); i += batch_size {
			max := int(math.Min(float64(i+batch_size), float64(len(data))))
			rides := bike_app_be.NewRidesFromDataSlice(keys, data[i:max])
			wg.Add(1)

		finished_goroutines_check:
			for j := 0; j < active_routines; j++ {
				select {
				case /*msg :=*/ <-consecutives_channel:
					//fmt.Printf("Routine %d finished\n", msg)
					active_routines--
				default:
					break finished_goroutines_check
				}
			}

			// if max consecutive, wait for one to finish
			if active_routines >= max_consecutive {
				<-consecutives_channel
				//fmt.Printf("Routine %d finished - max\n", msg)
				active_routines--
			}

			active_routines++
			i := i
			go func(done chan int, rides []*bike_app_be.Ride) {
				defer wg.Done()
				err := ride_service.CreateRides(rides)
				if err != nil {
					fmt.Println(err.Error())
				}
				consecutives_channel <- i
				//fmt.Println(i)
			}(consecutives_channel, rides)
			print_update(i)
		}
		wg.Wait()
		fmt.Println()

	}
	return nil
}
