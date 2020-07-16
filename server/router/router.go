package router

import (
	"../middleware"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/add", middleware.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/index", middleware.IndexAll).Methods("GET", "OPTIONS")

	return router
}
