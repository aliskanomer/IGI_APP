package utils

import (
	// modules
	"encoding/json"
	"net/http"
)

// Parses JSON response in given interface
func ParseJSONResponse(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()               // release resources
	decoder := json.NewDecoder(resp.Body) // create a new decoder
	return decoder.Decode(target)         // decode the response body into the target interface
}
