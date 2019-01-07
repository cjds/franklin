// Author(s): Carl Saldanha

package server

import (
	"net/http"

	"github.com/gorilla/mux"
		  "github.com/jinzhu/gorm"
		    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

// LoadRoutes will return a Gorilla mux router as a http handler for request routing.
func LoadRoutes(db *gorm.DB) http.Handler {
	// Instantiate router
	muxRouter := mux.NewRouter().StrictSlash(true)
	//muxRouter.Use(
	//	NewCORSMiddleware(),
	//	NewServerLoggingMiddleware(),
	//	NewAuthMiddleware(),
	//	NewGzipMiddleware(),
	//)

	// Namespacing API
	api := muxRouter.PathPrefix("/franklin/api/v1").Subrouter()
	api.Handle("/", NewVersionHandler()).Methods("GET")

	// APIs
	handlePostTask := NewPostTaskHandler(db)
	api.Handle("/tasks", handlePostTask).Methods("POST")
	api.Handle("/tasks", NewGetTaskHandler(db)).Methods("GET")



	return muxRouter
}
