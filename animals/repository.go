package animals

import (
	"database/sql"
)

type AnimalRepository struct {
	DB *sql.DB
}

func NewAnimalRepository(db *sql.DB) *AnimalRepository {
	return &AnimalRepository{DB: db}
}

// Create a new animal
func (r *AnimalRepository) Create(animal *Animal) error {
	sqlStatement := `INSERT INTO animals (id, name, class, legs) VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(sqlStatement, animal.ID, animal.Name, animal.Class, animal.Legs)
	return err
}

// Get all animals
func (r *AnimalRepository) GetAll() ([]Animal, error) {
	var animals []Animal
	rows, err := r.DB.Query("SELECT * FROM animals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var animal Animal
		err := rows.Scan(&animal.ID, &animal.Name, &animal.Class, &animal.Legs, &animal.CreatedAt)
		if err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	return animals, nil
}

// Update animal by ID
func (r *AnimalRepository) Update(id int, animal *Animal) error {
	sqlStatement := `UPDATE animals SET name=$1, Class=$2, Legs=$3 WHERE id=$4`
	_, err := r.DB.Exec(sqlStatement, animal.Name, animal.Class, animal.Legs, id)
	return err
}

// Delete animal by ID
func (r *AnimalRepository) Delete(id int) error {
	sqlStatement := `DELETE FROM animals WHERE id=$1`
	_, err := r.DB.Exec(sqlStatement, id)
	return err
}

// GetByID fetches an animal by its ID
func (r *AnimalRepository) GetByID(id int) (*Animal, error) {
	var animal Animal
	err := r.DB.QueryRow("SELECT id, name, class, legs FROM animals WHERE id = $1", id).Scan(
		&animal.ID, &animal.Name, &animal.Class, &animal.Legs,
	)
	if err == sql.ErrNoRows {
		return nil, nil // Return nil if no animal found
	}
	if err != nil {
		return nil, err
	}
	return &animal, nil
}
