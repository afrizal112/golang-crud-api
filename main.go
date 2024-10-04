package main

import (
	"log"
	"net/http"

	"github.com/yourusername/golang-crud-api/database"
	"github.com/yourusername/golang-crud-api/pkg/router"
)

func main() {
	// Initialize the database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize the router
	r := router.InitRouter(db)

	log.Println("API is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
