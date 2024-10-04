package router

import (
	"database/sql"

	"github.com/afrizal112/golang-crud-api/animals"
	"github.com/gorilla/mux"
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
