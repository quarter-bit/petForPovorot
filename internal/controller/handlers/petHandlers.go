package handlers

import (
	"Pet_store/internal/models"
	"log"

	"encoding/json"
	"net/http"
	"strconv"
)

type PetHandleIFace interface {
	AddPet(pet models.Pet) error
	FullStructPetUpdate(pet models.Pet) error
	FindPetByStatus(status string) (models.Pet, error)
	//GetByTags(w http.ResponseWriter, r *http.Request)
	FindPetById(id int) (models.Pet, error)
	UpdatePetByForm(Id int, name string, status string) error
	DeletePetById(id int) error
}

type PetHandlers struct {
	Service PetHandleIFace
}

func NewPetHandlers(service PetHandleIFace) *PetHandlers {
	return &PetHandlers{
		Service: service,
	}
}

// @Summary Add pet
// @Security ApiKeyAuth
// @Tags Pet
// @Accept json
// @Produce json
// @Param pet body models.Pet true "Pet"
// Success 200 {object} models.Pet
// @Router /pet [post]
func (h *PetHandlers) Add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pet models.Pet
		err := json.NewDecoder(r.Body).Decode(&pet)
		defer r.Body.Close()
		if err != nil {
			log.Fatal("Ошибка чтения тела запроса: ", err, "\n", pet)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = h.Service.AddPet(pet)
		if err != nil {
			log.Fatal("Ошибка добавления питомца: ", err, "\n", pet)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println("Питомец добавлен: ", pet)
		w.WriteHeader(http.StatusOK)
	}
}

// @Summary Find pet by ID
// @Security ApiKeyAuth
// @Tags Pet
// @Accept json
// @Produce json
// @Param id query string true "ID"
// Success 200 {object} models.Pet
// @Router /pet/{petId} [get]
func (h *PetHandlers) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println("Ошибка чтения тела запроса: ", err, "\n", idStr)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		pet, err := h.Service.FindPetById(id)
		if err != nil {
			log.Println("Ошибка поиска питомца: ", err, "\n", pet)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println("Питомец найден: ", pet)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(pet.ToJson()))
	}

}

// @Summary Find pets by status
// @deprecated
// @Security ApiKeyAuth
// @Tags Pet
// @Accept json
// @Produce json
// @Param status query string false "Status"
// Success 200 {object} models.Pet
// @Router /pet/findByStatus [get]
func (h *PetHandlers) GetByStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := r.URL.Query().Get("status")
		pets, err := h.Service.FindPetByStatus(status)
		if err != nil {
			log.Println("Error getting pets by status: ", err, "\n", pets)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println("Pets by status found: ", pets)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(pets.ToJson()))
	}

}

// func (h *PetHandlers) GetByTag(w http.ResponseWriter, r *http.Request) {

// 	tag := r.URL.Query().Get("tag")
// 	pets, err := h.service.FindPetByTag(tag)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(pets.ToJson()))
// }

// @Summary Delete pet by ID
// @Security ApiKeyAuth
// @Tags Pet
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Router /pet/{petId} [delete]
func (h *PetHandlers) DeleteById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println("Error reading id: ", err, "\n", idStr)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = h.Service.DeletePetById(id)
		if err != nil {
			log.Println("Error deleting pet: ", err, "\n", id)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println("Pet deleted: ", id)
		w.WriteHeader(http.StatusOK)
	}

}

// @Summary Update pet by form
// @Security ApiKeyAuth
// @Tags Pet
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Param name query string false "Name"
// @Param status query string false "Status"
// @Router /pet/{petId} [post]
func (h *PetHandlers) FormUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println("Ошибка чтения id: ", err, "\n", idStr)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		name := r.URL.Query().Get("name")

		status := r.URL.Query().Get("status")

		// var pet models.Pet
		// err = json.NewDecoder(r.Body).Decode(&pet)
		// if err != nil {
		// 	log.Fatal("Ошибка чтения тела запроса: ", err, "\n", pet)
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }
		err = h.Service.UpdatePetByForm(id, name, status)
		if err != nil {
			log.Println("Ошибка обновления питомца: ", err, "\n", id, "\n", name, "\n", status)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println("Питомец обновлен: ", id, "\n", name, "\n", status)
		w.WriteHeader(http.StatusOK)
	}

}

// @Summary Update pet by JSON
// @Security ApiKeyAuth
// @Tags Pet
// @Accept json
// @Produce json
// @Param pet body models.Pet true "Pet"
// @Router /pet [put]
func (h *PetHandlers) FullUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pet models.Pet
		err := json.NewDecoder(r.Body).Decode(&pet)
		if err != nil {
			log.Fatal("Ошибка чтения тела запроса: ", err, "\n", pet)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = h.Service.FullStructPetUpdate(pet)
		if err != nil {
			log.Fatal("Ошибка обновления питомца: ", err, "\n", pet)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println("Питомец обновлен: ", pet)
		w.WriteHeader(http.StatusOK)
	}

}
