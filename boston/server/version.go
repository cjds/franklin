// Author(s): Carl Saldanha

package server

import (
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"
)

// NewVersionHandler handles version GET requests.
func NewVersionHandler() http.HandlerFunc {
	info := struct {
		Version string `json:"version"`
	}{
		Version: viper.GetString("version"),
	}

	infoBytes, _ := json.Marshal(info)

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(infoBytes)
	}
}
