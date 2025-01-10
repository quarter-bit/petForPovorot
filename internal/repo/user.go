package repo

import (
	"Pet_store/internal/models"
	"database/sql"
	"log"
	"reflect"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) GetByName(name string) (models.User, error) {
	var user models.User
	r.db.QueryRow("SELECT id, username, first_name, last_name, email, password, phone, user_status FROM users WHERE username = $1", name).Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Phone, &user.UserStatus)
	if reflect.DeepEqual(user, models.User{}) {
		log.Fatal(user)
		return user, sql.ErrNoRows
	}
	return user, nil
}

func (r *UserRepo) Update(user models.User) error {
	_, err := r.db.Exec("UPDATE users SET username = $1, first_name = $2, last_name = $3, email = $4, password = $5, phone = $6, user_status = $7, WHERE id = $8", user.Username, user.Email, user.Password, user.UserStatus, user.ID)
	return err
}

func (r *UserRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func (r *UserRepo) Create(user models.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, username, first_name, last_name, email, password, phone, user_status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", user.ID, user.Username, user.FirstName, user.LastName, user.Email, user.Password, user.Phone, user.UserStatus)
	return err
}

func (r *UserRepo) CreateWithGivenInputArray(users []models.User) ([]models.User, error) {
	var result []models.User
	for _, user := range users {
		err := r.Create(user) //db.Exec("INSERT INTO users (id, username, email, password, user_status) VALUES (?, ?, ?, ?, ?)", user.ID, user.Username, user.Email, user.Password, user.UserStatus)
		if err != nil {
			return nil, err
		}
		result = append(result, user)
	}
	return result, nil
}
