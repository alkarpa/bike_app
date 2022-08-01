package database

import (
	"database/sql"
	"fmt"
	"math"
	"os"
	"strings"
	"sync"

	"alkarpa.fi/bike_app_be"
	csv "alkarpa.fi/csv_reader"
)

type CSVImporter struct {
	Path     string
	Verbose  bool
	Database *sql.DB
}

func (impo *CSVImporter) printLn(msg string) {
	if impo.Verbose {
		fmt.Println(msg)
	}
}
func (impo *CSVImporter) print(msg string) {
	if impo.Verbose {
		fmt.Print(msg)
	}
}

// Initial placeholder for testing and base for future modifications
func (impo *CSVImporter) ImportFromCSVs() error {
	impo.printLn("Import from CSV")

	/*
		db, err := OpenSQL()
		if err != nil {
			return err
		}
	*/

	impo.printLn("\n-=Stations")
	station_ids, err := impo.csvStations(impo.Database)
	if err != nil {
		return err
	}

	impo.printLn("\n-=Rides")
	if err = impo.csvRides(impo.Database, station_ids); err != nil {
		return err
	}

	return nil
}

func (impo *CSVImporter) csvStations(db *sql.DB) (map[int]struct{}, error) {
	directory := impo.Path + "station/" // "../data/station/"
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}
	station_ids := make(map[int]struct{})
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".csv") {
			continue
		}
		path := fmt.Sprintf("%s%s", directory, f.Name())
		impo.printLn(fmt.Sprintf("Importing %s\n", path))
		stids, err := impo.addStationsFromCsv(db, path)
		if err != nil {
			fmt.Println(err.Error())
		}
		for key, value := range stids {
			station_ids[key] = value
		}
	}
	return station_ids, nil
}
func (impo *CSVImporter) addStationsFromCsv(db *sql.DB, path string) (map[int]struct{}, error) {
	keys, data, err := csv.ReadFromCSV(path)
	if err != nil {
		return nil, err
	}

	stations := bike_app_be.NewStationsFromDataSlice(keys, data)

	station_service := NewStationService(db)
	if err = station_service.InsertStations(stations); err != nil {
		return nil, err
	}
	station_id_set := make(map[int]struct{})
	for _, station := range stations {
		station_id_set[station.Id] = struct{}{}
	}
	return station_id_set, station_service.insertStationLangFields(stations)
}

func (impo *CSVImporter) csvRides(db *sql.DB, station_ids map[int]struct{}) error {
	directory := impo.Path + "ride/" // "../data/ride/"
	files, err := os.ReadDir(directory)
	if err != nil {
		return err
	}
	ride_file_names := make([]string, 0, len(files))
	for _, f := range files {
		ride_file_names = append(ride_file_names, fmt.Sprintf("%s%s", directory, f.Name()))
	}

	ride_service := NewRideService(db)
	for _, path := range ride_file_names {
		import_info := fmt.Sprintf("Importing '%s'", path)
		impo.printLn(import_info)
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
			impo.print(fmt.Sprintf("\r%s - [%7d / %7d]", import_info, i, len(data)))
		}

		for i := 0; i < len(data); i += batch_size {
			max := int(math.Min(float64(i+batch_size), float64(len(data))))
			rides := bike_app_be.NewRidesFromDataSlice(station_ids, keys, data[i:max])
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
		print_update(len(data))
		impo.printLn("")

	}
	return nil
}
