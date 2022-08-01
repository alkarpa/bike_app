package main

import (
	"log"

	"alkarpa.fi/bike_app_be/database"
	"alkarpa.fi/bike_app_be/http"
)

func main() {
	db, err := database.OpenSQL()
	if err != nil {
		log.Fatal("failed to open database connection")
	}
	defer db.Close()

	server := http.NewServer()
	server.RideService = database.NewRideService(db)
	server.StationService = database.NewStationService(db)

	server.ListenAndServe()
}
