package plivogo

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server

	//client is Plivo client being tested
	client *Client
)

const (
	testAuthID    = "MANJQZNJLJZWY3ZJK5ZW"
	testAuthToken = "5678"
)

func testSetup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(testAuthID, testAuthToken)
	client.baseURL = server.URL
}

// teardown closes the test HTTP server.
func testTeardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if want != r.Method {
		t.Errorf("Request method = %v, want %v", r.Method, want)
	}
}
