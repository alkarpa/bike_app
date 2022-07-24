package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) getStations() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		stations, err := server.StationService.GetAll()
		if err != nil {
			http.Error(w, "error retrieving stations", http.StatusInternalServerError)
			return
		}

		js, _ := json.Marshal(stations)
		w.Write(js)

	}
}

func (server *Server) getStationDetails() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		station_id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, "error reading id", http.StatusBadRequest)
			return
		}
		station, err := server.StationService.GetDetails(station_id)
		if err != nil {
			http.Error(w, "error retrieving station", http.StatusInternalServerError)
			return
		}
		js, _ := json.Marshal(station)
		w.Write(js)
	}
}
