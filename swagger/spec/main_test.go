package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {

	// Prepare GET request to "/" endpoint
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error("could to prepare request: ", err)
	}

	// record the response
	rr := httptest.NewRecorder()

	// Start test server prepared to receive request on 'ping' endpoint
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(rr, req)

	// We expect to get 404. Really.
	// Remember that swagger.yaml will be copied to this directory
	// when building the Docker image.
	expectedStatus := 404

	if rr.Code != expectedStatus {
		t.Errorf("expected status: %d, observed: %d", rr.Code, expectedStatus)
	}
}
