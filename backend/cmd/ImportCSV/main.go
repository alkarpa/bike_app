package main

import (
	"fmt"

	db "alkarpa.fi/bike_app_be/database"
)

// For manual csv import testing; possibly a manual csv import tool later
func main() {
	if err := db.ImportFromCSVs(); err != nil {
		fmt.Println(err.Error())
	}
}
