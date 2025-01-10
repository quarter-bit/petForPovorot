package ifaces

import "Pet_store/internal/models"

// upload foto
// add new pet to the store
// update pet
// finde by status
// find by tag
// find bi id
// update by form
// delite by id
type PetServiceIFace interface {
	AddPet(pet models.Pet) error
	UpdatePetByFullStruct(pet models.Pet) error
	FindPetByStatus(status string) (models.Pet, error)
	// FindPetByTag(tag string) ([]models.Pet, error)
	FindPetById(id int) (models.Pet, error)
	UpdatePetByForm(Id int, name string, status string) error
	DeletePetById(id int) error
}
