package service

import (
	"Pet_store/internal/models"
	"Pet_store/internal/service/ifaces"
)

type StoreService struct {
	Repo ifaces.StoreServiceIFace
}

func NewStoreService(repo ifaces.StoreServiceIFace) *StoreService {
	return &StoreService{
		Repo: repo,
	}
}

func (s *StoreService) Inventory() []models.Pet {
	return s.Repo.Inventory()
}

func (s *StoreService) CreateOrder(order models.Order) error {
	return s.Repo.CreateOrder(order)
}

func (s *StoreService) GetOrderById(orderId int) (models.Order, error) {
	return s.Repo.GetOrderById(orderId)
}

func (s *StoreService) DeleteOrder(orderId int) error {
	return s.Repo.DeleteOrder(orderId)
}
