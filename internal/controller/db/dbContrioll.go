package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DBController struct {
	Db *sql.DB
}

func CloseDB(db *sql.DB) {
	db.Close()
}

func Connect() (*DBController, error) {
	// Загружаем переменные окружения из .env файла
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME")))
	if err != nil {
		return &DBController{}, err
	}

	// Проверяем соединение с базой данных
	err = db.Ping()
	if err != nil {
		return &DBController{}, fmt.Errorf("Error pinging the database: %v", err)
	}

	// Создаем таблицу, если она не существует
	err = createTableIfNotExists(db)
	if err != nil {
		return &DBController{}, fmt.Errorf("Error creating table: %v", err)
	}
	return &DBController{db}, nil
}

func createTableIfNotExists(db *sql.DB) error {
	queryUsers := `
  CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	phone VARCHAR(255) NOT NULL,
	user_status INT NOT NULL
  );
  `
	queryPets := `
  CREATE TABLE IF NOT EXISTS pets (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	status VARCHAR(255) NOT NULL
  );
  `
	queryOrders := `
  CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    pet_id INT NOT NULL,
	quantity INT NOT NULL,
	ship_date VARCHAR(255) NOT NULL,	
	status VARCHAR(255) NOT NULL,
	complete BOOLEAN	
  );
  `
	queryCategories := `
  CREATE TABLE IF NOT EXISTS pet_categories (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL
  );`
	queryTags := `
  CREATE TABLE IF NOT EXISTS pets_tags (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL
  );`
	queryFotoUrls := `
  CREATE TABLE IF NOT EXISTS pets_foto_urls (
	id SERIAL PRIMARY KEY,
	url VARCHAR(255) NOT NULL
  )`

	_, err := db.Exec(queryCategories)
	if err != nil {
		return fmt.Errorf("Error creating table categories: %v", err)
	}
	_, err = db.Exec(queryTags)
	if err != nil {
		return fmt.Errorf("Error creating table tags: %v", err)
	}
	_, err = db.Exec(queryFotoUrls)
	if err != nil {
		return fmt.Errorf("Error creating table foto_urls: %v", err)
	}
	_, err = db.Exec(queryUsers)
	if err != nil {
		return fmt.Errorf("Error creating table users: %v", err)
	}
	_, err = db.Exec(queryPets)
	if err != nil {
		return fmt.Errorf("Error creating table pets: %v", err)
	}
	_, err = db.Exec(queryOrders)
	if err != nil {
		return fmt.Errorf("Error creating table orders: %v", err)
	}

	return err
}
