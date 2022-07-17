package http

import (
	"encoding/json"
	"net/http"
)

func (server *Server) getRides() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		rides, err := server.RideService.GetRides()
		if err != nil {
			http.Error(w, "error retrieving stations", http.StatusInternalServerError)
		}

		js, _ := json.Marshal(rides)
		w.Write(js)

	}
}
