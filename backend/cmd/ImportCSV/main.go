package main

import (
	"fmt"
	"time"

	db "alkarpa.fi/bike_app_be/database"
)

// For manual csv import testing; possibly a manual csv import tool later
func main() {
	start := time.Now()
	if err := db.ImportFromCSVs(); err != nil {
		fmt.Println(err.Error())
	}
	end := time.Now()
	diff := end.Sub(start)
	fmt.Printf("Took %s \n", diff)
}
