package server

import (
	"net"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServe(t *testing.T) {
	timeOut := time.Duration(3) * time.Second
	server := httptest.NewServer(serverEngine())
	defer server.Close()

	// fixes weird double ':' problem
	port := server.URL[len(server.URL)-5:]

	_, err := net.DialTimeout("tcp", "localhost:"+port, timeOut)
	if err != nil {
		t.Errorf("failed to dial server: %s", err)
	}
}

func TestApiVersion(t *testing.T) {
	expected := "v1"
	observed := apiVersion()
	if observed != expected {
		t.Errorf("observed: %s, expected: %s", observed, expected)
	}
}
