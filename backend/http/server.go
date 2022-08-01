package http

import (
	"fmt"
	"net/http"

	"alkarpa.fi/bike_app_be"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router

	RideService    bike_app_be.RideService
	StationService bike_app_be.StationService
}

func NewServer() *Server {

	server := &Server{
		router: mux.NewRouter(),
	}

	{
		subrouter := server.router.PathPrefix("/ride").Subrouter()
		server.registerRideRoutes(subrouter)
	}
	{
		subrouter := server.router.PathPrefix("/station").Subrouter()
		server.registerStationRoutes(subrouter)
	}

	server.router.Use(mux.CORSMethodMiddleware(server.router))
	server.router.Use(middleware)

	return server
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func (server *Server) registerRideRoutes(r *mux.Router) {
	r.HandleFunc("/", server.getRides()).Methods("GET")
}
func (server *Server) registerStationRoutes(r *mux.Router) {
	r.HandleFunc("/", server.getStations()).Methods("GET")
	r.HandleFunc("/{id}", server.getStationDetails()).Methods("GET")
}

func (server *Server) ListenAndServe() {
	fmt.Println("Opening server")
	http.ListenAndServe("localhost:8080", server.router)

}
