package http

import (
	"testing"

	"alkarpa.fi/bike_app_be/mock"
)

type TestServer struct {
	*Server

	RideService mock.RideService
}

func OpenTestServer(tb testing.TB) *TestServer {
	tb.Helper()

	ts := &TestServer{Server: NewServer()}
	ts.Server.RideService = &ts.RideService

	ts.ListenAndServe()

	return ts
}
