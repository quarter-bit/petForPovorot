package ifaces

import (
	"Pet_store/internal/models"
)

//Get by name
//update
//delite
//create
//create with giwing input array
//login
//logout

type UserServiceIFace interface {
	GetByName(name string) (models.User, error)
	Update(user models.User) error
	Delete(id int) error
	Create(user models.User) error
	CreateWithGivenInputArray(users []models.User) ([]models.User, error)
}
