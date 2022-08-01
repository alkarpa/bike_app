package main

import (
	"fmt"
	"time"

	"alkarpa.fi/bike_app_be/database"
)

// For manual csv import testing; possibly a manual csv import tool later
func main() {
	start := time.Now()
	db, err := database.OpenSQL()
	if err != nil {
		fmt.Println(err)
	}
	importer := database.CSVImporter{
		Path:     "../data/",
		Verbose:  true,
		Database: db,
	}
	if err := importer.ImportFromCSVs(); err != nil {
		fmt.Println(err.Error())
	}
	end := time.Now()
	diff := end.Sub(start)
	fmt.Printf("Took %s \n", diff)
}
