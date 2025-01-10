package ifaces

import (
	"Pet_store/internal/models"
)

type StoreServiceIFace interface {
	Inventory() []models.Pet
	CreateOrder(order models.Order) error
	GetOrderById(orderId int) (models.Order, error)
	DeleteOrder(orderId int) error
}
