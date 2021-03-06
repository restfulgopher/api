package server

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidHandlerStatusCode(t *testing.T) {
	expectedStatus := 200

	server := httptest.NewServer(serverEngine())
	defer server.Close()

	resp, err := http.Get(server.URL + "/v1/iban/valid/123")
	if err != nil {
		t.Errorf("failed to execute GET request: %s", err)
	}

	if resp.StatusCode != expectedStatus {
		t.Errorf("expected status: %d, observed: %d", resp.StatusCode, expectedStatus)
	}
}

func TestValidResponseHandler(t *testing.T) {
	expectedResponses := []string{
		"{\"iban\":\"123\",\"valid\":false}", // zsh
		"{\"iban\":\"123\",\"valid\":false}\n",
	}

	server := httptest.NewServer(serverEngine())
	defer server.Close()

	resp, err := http.Get(server.URL + "/v1/iban/valid/123")
	if err != nil {
		t.Errorf("failed to execute GET request: %s", err)
	}
	defer resp.Body.Close()

	var b bytes.Buffer
	if _, err := io.Copy(&b, resp.Body); err != nil {
		t.Errorf("failed to copy response body: %s", err)
	}

	if expectedResponses[0] != b.String() && expectedResponses[1] != b.String() {
		t.Errorf("expected response: %q or %q, observed: %q",
			expectedResponses[0], expectedResponses[0], b.String())
	}
}
