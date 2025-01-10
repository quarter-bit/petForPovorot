package repo

import (
	"Pet_store/internal/models"
	"database/sql"
)

type PetRepo struct {
	db *sql.DB
}

func NewPetRepo(db *sql.DB) *PetRepo {
	return &PetRepo{
		db: db,
	}
}

func (r *PetRepo) AddPet(pet models.Pet) error {
	_, err := r.db.Exec("INSERT INTO pets (id, name, status) VALUES ($1, $2, $3)", pet.ID, pet.Name, pet.Status)
	if err != nil {
		return err
	}
	for _, tag := range pet.Tags {
		_, err = r.db.Exec("INSERT INTO pets_tags (id, name) VALUES ($1, $2)", tag.ID, tag.Name)
		if err != nil {
			return err
		}
	}
	for _, photo := range pet.PhotoUrls {
		_, err = r.db.Exec("INSERT INTO pets_foto_urls (id, url) VALUES ($1, $2)", pet.ID, photo)
		if err != nil {
			return err
		}
	}
	_, err = r.db.Exec("INSERT INTO pet_categories (id, name) VALUES ($1, $2)", pet.Category.ID, pet.Category.Name)
	return err
}

func (r *PetRepo) UpdatePetByFullStruct(pet models.Pet) error {
	_, err := r.db.Exec("UPDATE pets SET name = $1, status = $2 WHERE id = $3", pet.Name, pet.Status, pet.ID)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("DELETE FROM pets_tags WHERE id = $1", pet.ID)
	if err != nil {
		return err
	}

	for _, tag := range pet.Tags {
		_, err = r.db.Exec("INSERT INTO pets_tags (id, name) VALUES ($1, $2)", tag.ID, tag.Name)
		if err != nil {
			return err
		}
	}

	_, err = r.db.Exec("DELETE FROM pets_foto_urls WHERE id = $1", pet.ID)
	if err != nil {
		return err
	}

	for _, photo := range pet.PhotoUrls {
		_, err = r.db.Exec("INSERT INTO pets_foto_urls (id, url) VALUES ($1, $2)", pet.ID, photo)
		if err != nil {
			return err
		}
	}

	_, err = r.db.Exec("UPDATE pet_categories SET name = $1 WHERE id = $2", pet.Category.Name, pet.Category.ID)
	if err != nil {
		return err
	}

	return nil

}

func (r *PetRepo) FindPetByStatus(status string) (models.Pet, error) {
	var pet models.Pet
	err := r.db.QueryRow("SELECT id, name FROM pets WHERE status = $1", status).Scan(&pet.ID, &pet.Name, &pet.Status)
	if err != nil {
		return models.Pet{}, err
	}

	err = r.db.QueryRow("SELECT id, name FROM pet_categories WHERE id = $1", pet.ID).Scan(&pet.Category.ID, &pet.Category.Name)
	if err != nil {
		return models.Pet{}, err
	}

	err = r.db.QueryRow("SELECT * FROM pets_tags WHERE id = $1", pet.ID).Scan(&pet.Tags)
	if err != nil {
		return models.Pet{}, err
	}

	err = r.db.QueryRow("SELECT * FROM pets_foto_urls WHERE id = $1", pet.ID).Scan(&pet.PhotoUrls)
	if err != nil {
		return models.Pet{}, err
	}

	return pet, nil
}

func (r *PetRepo) FindPetById(id int) (models.Pet, error) {
	return Finder(r.db, id)
}

// func (r *PetRepo) FindPetByTag(tag string) ([]models.Pet, error) {
// 	pets := make([]models.Pet, 0)
// 	//ids := make([]int, 0)

// 	raws, err := r.db.Query("SELECT id FROM pets_tags WHERE name = ?", tag)
// 	if err != nil {
// 		return pets, err
// 	}

// 	for raws.Next() {
// 		var id int
// 		err := raws.Scan(&id)
// 		if err != nil {
// 			return pets, err
// 		}

// 		var pet models.Pet
// 		err = r.db.QueryRow("SELECT * FROM pets WHERE id = ?", id).Scan(&pet.ID, &pet.Name, &pet.Status)
// 		if err != nil {
// 			return pets, err
// 		}
// 		err = r.db.QueryRow("SELECT * FROM pets_categories WHERE id = ?", pet.ID).Scan(&pet.Category.ID, &pet.Category.Name)
// 		if err != nil {
// 			return pets, err
// 		}
// 		err = r.db.QueryRow("SELECT * FROM pets_tags WHERE pet_id = ?", pet.ID).Scan(&pet.Tags)
// 		if err != nil {
// 			return pets, err
// 		}
// 		err = r.db.QueryRow("SELECT * FROM pets_foto_urls WHERE id = ?", pet.ID).Scan(&pet.PhotoUrls)
// 		if err != nil {
// 			return pets, err
// 		}

// 		pets = append(pets, pet)
// 	}

// 	return pets, nil
// }

func (r *PetRepo) UpdatePetByForm(Id int, name string, status string) error {
	_, err := r.db.Exec("UPDATE pets SET name = $1, status = $2 WHERE id = $3", name, status, Id)
	return err
}

func (r *PetRepo) DeletePetById(id int) error {
	_, err := r.db.Exec("DELETE FROM pets WHERE id = $1", id)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("DELETE FROM pet_categories WHERE id = $1", id)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("DELETE FROM pets_tags WHERE id = $1", id)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("DELETE FROM pets_foto_urls WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
