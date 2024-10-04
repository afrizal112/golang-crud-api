package animals

type AnimalService struct {
	Repo *AnimalRepository
}

func NewAnimalService(repo *AnimalRepository) *AnimalService {
	return &AnimalService{Repo: repo}
}

func (s *AnimalService) CreateAnimal(animal *Animal) error {
	return s.Repo.Create(animal)
}

func (s *AnimalService) GetAllAnimals() ([]Animal, error) {
	return s.Repo.GetAll()
}

func (s *AnimalService) GetAnimalByID(id int) (*Animal, error) {
	return s.Repo.GetByID(id)
}

func (s *AnimalService) UpdateAnimal(id int, animal *Animal) error {
	return s.Repo.Update(id, animal)
}

func (s *AnimalService) DeleteAnimal(id int) error {
	return s.Repo.Delete(id)
}
