package routes

import (
	"auth/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	// Authentication Routes
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
}
