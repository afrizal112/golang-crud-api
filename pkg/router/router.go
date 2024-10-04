package router

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/yourusername/golang-crud-api/animals"
)

func InitRouter(db *sql.DB) *mux.Router {
	repo := animals.NewAnimalRepository(db)
	service := animals.NewAnimalService(repo)
	handler := animals.NewAnimalHandler(service)

	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/animal", handler.CreateAnimal).Methods("POST")
	r.HandleFunc("/animals", handler.GetAnimals).Methods("GET")
	r.HandleFunc("/animal/{id}", handler.GetAnimalByID).Methods("GET")
	r.HandleFunc("/animal/{id}", handler.UpdateAnimal).Methods("PUT")
	r.HandleFunc("/animal/{id}", handler.DeleteAnimal).Methods("DELETE")

	return r
}
