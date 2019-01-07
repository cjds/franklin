// Author(s): Carl Saldanha

package server

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse is the default json returned when an error has occured.
// For example,
// {
//   "status": 404,
//   "message": "record not found"
// }
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// writeError writes an error response to response writer.
func writeError(w http.ResponseWriter, code int, msg string) {
	res := &ErrorResponse{
		Status:  code,
		Message: msg,
	}

	if bytes, err := json.Marshal(res); err == nil {
		w.WriteHeader(code)
		w.Write(bytes)
	}
}
