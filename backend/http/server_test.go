package http

import (
	"testing"

	"alkarpa.fi/bike_app_be/mock"
)

type TestServer struct {
	*Server

	RideService    mock.RideService
	StationService mock.StationService
}

func OpenTestServer(tb testing.TB) *TestServer {
	tb.Helper()

	ts := &TestServer{Server: NewServer()}
	ts.Server.RideService = &ts.RideService
	ts.Server.StationService = &ts.StationService

	return ts
}
