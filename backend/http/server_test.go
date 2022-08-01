package http

import (
	"fmt"
	"testing"

	"alkarpa.fi/bike_app_be/mock"
)

const test_port = "8080"

var test_url = fmt.Sprintf("http://%s:%s", addr, test_port)

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

	ts.ListenAndServe()

	return ts
}
