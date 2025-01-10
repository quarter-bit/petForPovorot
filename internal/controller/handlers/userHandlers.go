package handlers

import (
	"Pet_store/internal/controller"
	"Pet_store/internal/models"
	"log"

	"encoding/json"
	"net/http"
	"strconv"
)

type UserControlIFace interface {
	// Login(w http.ResponseWriter, r *http.Request)
	// Logout(w http.ResponseWriter, r *http.Request)
	GetByName(name string) (models.User, error)
	Update(user models.User) error
	Delete(id int) error
	Create(user models.User) error
	CreateWithGivenInputArray(users []models.User) ([]models.User, error)
}

type UserHandlers struct {
	Service UserControlIFace
}

func NewUserHandlers(service UserControlIFace) *UserHandlers {
	return &UserHandlers{
		Service: service,
	}
}

// @Summary Login User
// @Description Login User
// @Tags user
// @Param username query string true "username"
// @Param password query string true "password"
// @Accept json
// @Produce json
// @Router /user/login [get]
// @Success 200 {string} string "User logged in successfully"
func (h *UserHandlers) Login( /*username, password string*/ ) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		password := r.URL.Query().Get("password")

		dbUser, err := h.Service.GetByName(username)
		log.Println(dbUser)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if password != dbUser.Password {
			log.Fatal("invalid password")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User logged in successfully\n"))
		token, _ := controller.GenerateJWToken(username)
		w.Write([]byte("Bearer "))
		w.Write([]byte(token))
	}

}

// @Summary Logout User
// @Description Logout User
// @Tags user
// @Accept json
// @Produce json
// @Router /user/logout [get]
// @Success 200 {string} string "logout"
func (h *UserHandlers) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text")
		w.Write([]byte("logout"))
	}
}

// @Summary Get User By Username
// @Description Get User By Username
// @Tags user
// @Param name query string true "username"
// @Accept json
// @Produce json
// @Router /user/{username} [get]
// @Success 200 {object} models.User
func (h *UserHandlers) GetByUsername() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		user, err := h.Service.GetByName(name)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(user.ToJson()))
	}

}

// @Summary Update User
// @Description Update User
// @Tags user
// @Accept json
// @Produce json
// @Router /user/{username} [put]
// @Success 200
func (h *UserHandlers) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = h.Service.Update(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

}

// @Summary Delete User
// @Description Delete User
// @Tags user
// @Accept json
// @Produce json
// @Param id query int true "ID"
// @Router /user/{username} [delete]
// @Success 200
func (h *UserHandlers) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = h.Service.Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

}

// @Summary Create User
// @Description Create User
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Router /user [post]
// @Success 200
func (h *UserHandlers) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = h.Service.Create(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

}

// @Summary Create User With Given Input Array
// @Description Create User
// @Tags user
// @Accept json
// @Produce json
// @Router /user/createWithArray [post]
// @Success 200
func (h *UserHandlers) CreateWithInputArray() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []models.User

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&users)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = h.Service.CreateWithGivenInputArray(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

}
