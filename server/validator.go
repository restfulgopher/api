package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// ValidationResponse wraps the result of the IBAN validation
type ValidationResponse struct {
	Iban  string `json:"iban"`
	Valid bool   `json:"valid"`
}

// validate iban using https://openiban.com/ api
func validate(iban string) (ValidationResponse, error) {
	URL := fmt.Sprintf("https://openiban.com/validate/%s", iban)

	// Prepare GET request
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return ValidationResponse{}, fmt.Errorf("could to create request: %s", err)
	}
	req.Header.Add("Accept", "application/json")

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ValidationResponse{}, fmt.Errorf("could not process get request to openiban api: %s", err)
	}
	defer resp.Body.Close()

	result, err := decode(resp.Body)
	if err != nil {
		return ValidationResponse{}, fmt.Errorf("could not validate iban: %s", err)
	}
	return result, nil
}

// decode response data into ValidationResponse struct
func decode(data io.Reader) (ValidationResponse, error) {
	var validationResp ValidationResponse

	dec := json.NewDecoder(data)
	if err := dec.Decode(&validationResp); err != nil {
		return ValidationResponse{}, fmt.Errorf("could not decode response from openiban to validation result map: %s", err)
	}
	return validationResp, nil
}

// sanitize remove any non-alphanumeric character from string
func sanitize(iban string) (string, error) {
	// Make a Regex to say we only want
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", fmt.Errorf("could not compile regexp expression: %s", err)
	}
	return reg.ReplaceAllString(iban, ""), nil
}
