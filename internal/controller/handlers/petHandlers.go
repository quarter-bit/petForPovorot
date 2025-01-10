package handlers

import (
	"Pet_store/internal/models"

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
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = h.Service.AddPet(pet)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

// @Summary Find pet by ID
// @Security ApiKeyAuth
// @Tags Pet
// @Accept json
// @Produce json
// @Param id path int true "ID"
// Success 200 {object} models.Pet
// @Router /pet/{petId} [get]
func (h *PetHandlers) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		pet, err := h.Service.FindPetById(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
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
			w.WriteHeader(http.StatusBadRequest)
			return
		}
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
// @Param id path int true "ID"
// @Router /pet/{petId} [delete]
func (h *PetHandlers) DeleteById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = h.Service.DeletePetById(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

}

// @Summary Update pet by form
// @Security ApiKeyAuth
// @Tags Pet
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param name query string false "Name"
// @Param status query string false "Status"
// @Router /pet/{petId} [post]
func (h *PetHandlers) FormUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		name := r.URL.Query().Get("name")
		// if err != nil {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }
		status := r.URL.Query().Get("status")
		// if err != nil {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }
		var pet models.Pet
		err = json.NewDecoder(r.Body).Decode(&pet)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = h.Service.UpdatePetByForm(id, name, status)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
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
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = h.Service.FullStructPetUpdate(pet)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

}
