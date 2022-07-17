package main

import (
	"context"
	"log"
	"os"
	"os/signal"

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

	// This makes the server wait for an interrupt
	ctx, cancel := context.WithCancel(context.Background())
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	go func() { <-channel; cancel() }()

	server.ListenAndServe()

	<-ctx.Done()

	server.Close()
}
