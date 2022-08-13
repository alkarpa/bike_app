package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"alkarpa.fi/bike_app_be"
)

func (server *Server) getRides() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		urlParams := r.URL.Query()
		//fmt.Printf("Query - %s", urlParams)

		count, err := server.RideService.GetCount(urlParams)
		if err != nil {
			http.Error(w, "error retrieving ride count", http.StatusInternalServerError)
		}

		rides, err := server.RideService.GetRides(urlParams)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "error retrieving rides", http.StatusInternalServerError)
		}
		response_data := &struct {
			Count int                 `json:"count"`
			Rides []*bike_app_be.Ride `json:"rides"`
		}{
			Count: count,
			Rides: rides,
		}

		js, _ := json.Marshal(response_data)
		w.Write(js)

	}
}
