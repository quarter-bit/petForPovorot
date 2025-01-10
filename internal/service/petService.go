package service

import (
	"Pet_store/internal/models"
	"Pet_store/internal/service/ifaces"
)

type PetService struct {
	Repo ifaces.PetServiceIFace
}

func NewPetService(repo ifaces.PetServiceIFace) *PetService {
	return &PetService{
		Repo: repo,
	}
}

func (s *PetService) AddPet(pet models.Pet) error {
	return s.Repo.AddPet(pet)
}

func (s *PetService) FullStructPetUpdate(pet models.Pet) error {
	return s.Repo.UpdatePetByFullStruct(pet)
}

func (s *PetService) FindPetByStatus(status string) (models.Pet, error) {
	return s.Repo.FindPetByStatus(status)
}

// func (s *PetService) FindPetByTag(tag string) (models.Pet, error) {
// 	return s.Repo.FindPetByTag(tag)
// }

func (s *PetService) FindPetById(id int) (models.Pet, error) {
	return s.Repo.FindPetById(id)
}

func (s *PetService) UpdatePetByForm(Id int, name string, status string) error {
	return s.Repo.UpdatePetByForm(Id, name, status)
}

func (s *PetService) DeletePetById(id int) error {
	return s.Repo.DeletePetById(id)
}
