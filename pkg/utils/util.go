package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// func to parse the request body into x. Returns nil if there are no errors.
func Parser(r *http.Request, x interface{}) []byte {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err1 := json.Unmarshal(body, x); err1 != nil {
			return ([]byte(err1.Error()))
		}
	}
	return nil
}
