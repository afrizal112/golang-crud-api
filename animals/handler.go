package animals

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AnimalHandler struct {
	Service *AnimalService
}

func NewAnimalHandler(service *AnimalService) *AnimalHandler {
	return &AnimalHandler{Service: service}
}

// Create a new animal
func (h *AnimalHandler) CreateAnimal(w http.ResponseWriter, r *http.Request) {
	var animal Animal
	if err := json.NewDecoder(r.Body).Decode(&animal); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if animal with the same ID already exists
	existingAnimal, err := h.Service.GetAnimalByID(animal.ID)
	if err != nil {
		http.Error(w, "Error checking for existing animal", http.StatusInternalServerError)
		return
	}
	if existingAnimal != nil {
		http.Error(w, "Animal with the specified ID already exists", http.StatusConflict)
		return
	}

	// If not exists, create the animal
	if err := h.Service.CreateAnimal(&animal); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(animal)
}

// Get all animals
func (h *AnimalHandler) GetAnimals(w http.ResponseWriter, r *http.Request) {
	animals, err := h.Service.GetAllAnimals()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(animals) == 0 {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animals)
}

// Get a specific animal by ID
func (h *AnimalHandler) GetAnimalByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid animal ID", http.StatusBadRequest)
		return
	}

	animal, err := h.Service.GetAnimalByID(id)
	if err != nil {
		http.Error(w, "Error fetching animal", http.StatusInternalServerError)
		return
	}
	if animal == nil {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animal)
}

// Update or create an animal if it doesn't exist
func (h *AnimalHandler) UpdateAnimal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid animal ID", http.StatusBadRequest)
		return
	}

	var animal Animal
	if err := json.NewDecoder(r.Body).Decode(&animal); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check if animal exists, update if found
	existingAnimal, err := h.Service.GetAnimalByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if existingAnimal == nil {
		// Create if not found
		animal.ID = id
		if err := h.Service.CreateAnimal(&animal); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(animal)
	} else {
		// Update if found
		if err := h.Service.UpdateAnimal(id, &animal); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(animal)
	}
}

// Delete an animal
func (h *AnimalHandler) DeleteAnimal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid animal ID", http.StatusBadRequest)
		return
	}

	// Check if animal exists
	existingAnimal, err := h.Service.GetAnimalByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if existingAnimal == nil {
		http.Error(w, "Animal not exist yet", http.StatusNotFound)
		return
	}

	// Proceed with deletion
	if err := h.Service.DeleteAnimal(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
