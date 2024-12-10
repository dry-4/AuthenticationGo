package main

import (
	"auth/config"
	"auth/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.GetDB()
	if err != nil {
		log.Fatal("Error getting db conneciotn")
	}

	if db != nil {
		sqlDB, _ := db.DB()
		defer sqlDB.Close()
	}

	// Setup router
	router := mux.NewRouter()
	routes.SetupRoutes(router)

	// start server
	port := ":8080"
	fmt.Printf("server starting at port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))

}
