package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"alkarpa.fi/bike_app_be"
	"github.com/gorilla/mux"
)

const addr = "localhost"
const port = "8080"

type Server struct {
	server *http.Server
	router *mux.Router

	RideService    bike_app_be.RideService
	StationService bike_app_be.StationService
}

func NewServer() *Server {

	server := &Server{
		server: &http.Server{},
		router: mux.NewRouter(),
	}

	server.server.Handler = http.HandlerFunc(server.serveHTTP)

	{
		subrouter := server.router.PathPrefix("/ride").Subrouter()
		server.registerRideRoutes(subrouter)
	}
	{
		subrouter := server.router.PathPrefix("/station").Subrouter()
		server.registerStationRoutes(subrouter)
	}

	return server
}

func (server *Server) registerRideRoutes(r *mux.Router) {
	r.HandleFunc("/", server.getRides()).Methods("GET")
}
func (server *Server) registerStationRoutes(r *mux.Router) {
	r.HandleFunc("/", server.getStations()).Methods("GET")
}

func (server *Server) ListenAndServe() {
	server.server.Addr = fmt.Sprintf("%s:%s", addr, port)
	go server.server.ListenAndServe()

	fmt.Println("Server listening")
}

func (server *Server) serveHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	server.router.ServeHTTP(w, r)
}

func (server *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.server.Shutdown(ctx)
	fmt.Println("Server shutdown")
}
