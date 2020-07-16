package router

import (
	"../middleware"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/planet/create", middleware.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/planet/index", middleware.IndexAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/planet/delete/{id}", middleware.DeleteByID).Methods("DELETE", "OPTIONS")

	return router
}
