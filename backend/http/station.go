package http

import (
	"encoding/json"
	"net/http"
)

func (server *Server) getStations() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		stations, err := server.StationService.GetAll()
		if err != nil {
			http.Error(w, "error retrieving stations", http.StatusInternalServerError)
		}

		js, _ := json.Marshal(stations)
		w.Write(js)

	}
}
