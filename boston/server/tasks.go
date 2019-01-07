// Author(s): Carl Saldanha

package server

import (
	"franklin/boston/model"
	"encoding/json"
	"github.com/jinzhu/gorm"
	  _ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"github.com/gofrs/uuid"
)

// NewPostHTTPMeasurementsHandler returns a http handler function that handles HTTPMeasurement
// POST requests.
func NewPostTaskHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		decoder := json.NewDecoder(r.Body)

		t := &model.Task{
		       ID:uuid.Must(uuid.NewV4()),
		}
		if err := decoder.Decode(t); err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := t.Insert(db); err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func NewGetTaskHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		tasks, err := model.GetTasks(db)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		t, err := json.Marshal(tasks)
		if err != nil {
			    // handle error
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
                w.Write(t)
	}
}
