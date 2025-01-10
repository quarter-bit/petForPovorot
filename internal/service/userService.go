package service

import (
	"Pet_store/internal/models"
	"Pet_store/internal/service/ifaces"
	//"Pet_store/internal/repository
)

type UserService struct {
	Repo ifaces.UserServiceIFace
}

func NewUserService(repo ifaces.UserServiceIFace) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (u *UserService) Login(userName, password string) ([]models.User, error) {
	return u.Login(userName, password)
}

func (u *UserService) GetByName(name string) (models.User, error) {
	return u.Repo.GetByName(name)
}

func (u *UserService) Update(user models.User) error {
	return u.Repo.Update(user)
}

func (u *UserService) Delete(id int) error {
	return u.Repo.Delete(id)
}

func (u *UserService) Create(user models.User) error {
	return u.Repo.Create(user)
}

func (u *UserService) CreateWithGivenInputArray(users []models.User) ([]models.User, error) {
	return u.Repo.CreateWithGivenInputArray(users)
}
