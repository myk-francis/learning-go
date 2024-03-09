package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API key from
// the headers of an HTTP request
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Malformed auth headers")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed first part of auth headers")
	}

	return vals[1], nil
}