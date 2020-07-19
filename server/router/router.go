package router

import (
	"../middleware"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/planet", middleware.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/planet", middleware.IndexAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/planet/{id}", middleware.IndexByID).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/planet/del/{id}", middleware.DeleteByID).Methods("DELETE", "OPTIONS")

	return router
}
