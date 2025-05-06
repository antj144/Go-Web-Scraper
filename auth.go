package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Get API key extracts an API Key
// the headers of an https request
// Example
// Authorization: APIKey instert apikey here
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.GET("Authorisation")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")

	}
	return vals[1], nil

}
