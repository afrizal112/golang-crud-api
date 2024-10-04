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
	sqlStatement := `INSERT INTO animals (name, species, age) VALUES ($1, $2, $3)`
	_, err := r.DB.Exec(sqlStatement, animal.Name, animal.Species, animal.Age)
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
		err := rows.Scan(&animal.ID, &animal.Name, &animal.Species, &animal.Age, &animal.CreatedAt)
		if err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	return animals, nil
}

// Update animal by ID
func (r *AnimalRepository) Update(id int, animal *Animal) error {
	sqlStatement := `UPDATE animals SET name=$1, species=$2, age=$3 WHERE id=$4`
	_, err := r.DB.Exec(sqlStatement, animal.Name, animal.Species, animal.Age, id)
	return err
}

// Delete animal by ID
func (r *AnimalRepository) Delete(id int) error {
	sqlStatement := `DELETE FROM animals WHERE id=$1`
	_, err := r.DB.Exec(sqlStatement, id)
	return err
}

func (r *AnimalRepository) GetByID(id int) (*Animal, error) {
	var animal Animal
	sqlStatement := `SELECT id, name, species, age, created_at FROM animals WHERE id=$1`
	row := r.DB.QueryRow(sqlStatement, id)

	err := row.Scan(&animal.ID, &animal.Name, &animal.Species, &animal.Age, &animal.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No animal found
		}
		return nil, err
	}
	return &animal, nil
}

// GetByName fetches an animal by its name
func (r *AnimalRepository) GetByName(name string) (*Animal, error) {
	var animal Animal
	err := r.DB.QueryRow("SELECT id, name, species, age FROM animals WHERE name = $1", name).Scan(
		&animal.ID, &animal.Name, &animal.Species, &animal.Age,
	)
	if err == sql.ErrNoRows {
		return nil, nil // Return nil if no animal found
	}
	if err != nil {
		return nil, err
	}
	return &animal, nil
}
