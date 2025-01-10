package repo

import (
	"Pet_store/internal/models"
	"database/sql"
)

type StoreRepo struct {
	db *sql.DB
}

func NewStoreRepo(db *sql.DB) *StoreRepo {
	return &StoreRepo{
		db: db,
	}
}

func Finder(db *sql.DB, id int) (models.Pet, error) {
	var pet models.Pet
	err := db.QueryRow("SELECT id, name, status FROM pets WHERE id = $1", id).Scan(&pet.ID, &pet.Name, &pet.Status)
	if err != nil {
		return models.Pet{}, err
	}

	err = db.QueryRow("SELECT id, name FROM pet_categories WHERE id = $1", pet.ID).Scan(&pet.Category.ID, &pet.Category.Name)
	if err != nil {
		return models.Pet{}, err
	}
	for i := range pet.Tags {
		err = db.QueryRow("SELECT id, name FROM pets_tags WHERE id = $1", pet.ID).Scan(&pet.Tags[i].ID, &pet.Tags[i].Name)
		if err != nil {
			return models.Pet{}, err
		}
	}
	for i := range pet.PhotoUrls {
		err = db.QueryRow("SELECT * FROM pets_foto_urls WHERE id = $1", pet.ID).Scan(&pet.PhotoUrls[i])
		if err != nil {
			return models.Pet{}, err
		}
	}

	return pet, nil
}

func (r *StoreRepo) Inventory() []models.Pet {
	pets := make([]models.Pet, 0)
	rows, err := r.db.Query("SELECT * FROM pets")
	if err != nil {
		return pets
	}

	for rows.Next() {
		var pet models.Pet
		err := rows.Scan(&pet.ID, &pet.Name, &pet.Status)
		if err != nil {
			return pets
		}
		pet, err = Finder(r.db, pet.ID)
		if err != nil {
			return pets
		}
		pets = append(pets, pet)
	}
	return pets
}

func (r *StoreRepo) CreateOrder(order models.Order) error {
	_, err := r.db.Exec("INSERT INTO orders (id, pet_id, quantity, ship_date, status, complete) VALUES ($1, $2, $3, $4, $5, $6)", order.ID, order.PetID, order.Quantity, order.ShipDate, order.Status, order.Complete)
	return err
}

func (r *StoreRepo) GetOrderById(orderId int) (models.Order, error) {
	var order models.Order
	err := r.db.QueryRow("SELECT * FROM orders WHERE id = $1", orderId).Scan(&order.ID, &order.PetID, &order.Quantity, &order.ShipDate, &order.Status, &order.Complete)
	return order, err
}

func (r *StoreRepo) DeleteOrder(orderId int) error {
	_, err := r.db.Exec("DELETE FROM orders WHERE id = $1", orderId)
	return err
}
