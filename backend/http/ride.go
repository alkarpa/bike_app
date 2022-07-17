package http

import (
	"encoding/json"
	"net/http"

	"alkarpa.fi/bike_app_be"
)

func (server *Server) getRides() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		placeholder := &bike_app_be.Ride{
			Departure:            "2022-07-10 12:37:00",
			Return:               "2022-07-10 12:38:00",
			Departure_station_id: 1,
			Return_station_id:    2,
			Distance:             100,
			Duration:             60,
		}

		js, err := json.Marshal(placeholder)
		if err != nil {
			return
		}
		w.Write(js)

	}
}
