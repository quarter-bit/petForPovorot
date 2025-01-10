package handlers

import (
	"Pet_store/internal/models"

	"encoding/json"
	"net/http"
	"strconv"
)

type StoreHandlerIFaces interface {
	Inventory() []models.Pet
	CreateOrder(order models.Order) error
	GetOrderById(orderId int) (models.Order, error)
	DeleteOrder(orderId int) error
}

type StoreHandler struct {
	Service StoreHandlerIFaces
}

func NewStoreHandler(service StoreHandlerIFaces) *StoreHandler {

	return &StoreHandler{
		Service: service,
	}
}

// @Summary Inventory
// @Security ApiKeyAuth
// @Tags Store
// @Accept json
// @Produce json
// @Success 200 {array} models.Pet
// @Router /store/inventory [get]
func (h *StoreHandler) InventoryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		respJson, err := json.Marshal(h.Service.Inventory())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write(respJson)
	}

}

// @Summary Create order
// @Tags Store
// @Accept json
// @Produce json
// @Param order body models.Order true "Order"
// @Success 200 {object} models.Order
// @Router /store/order [post]
func (h *StoreHandler) CreateOrderHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var order models.Order
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = h.Service.CreateOrder(order)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

}

// @Summary Get order by id
// @Tags Store
// @Accept json
// @Produce json
// @Param id path int true "Order id"
// @Success 200 {object} models.Order
// @Router /store/order/{orderId} [get]
func (h *StoreHandler) GetOrderByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		order, err := h.Service.GetOrderById(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		respJson, err := json.Marshal(order)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJson)
	}

}

// @Summary Delete order by id
// @Tags Store
// @Accept json
// @Produce json
// @Param id path int true "Order id"
// @Success 200
// @Router /store/order/{orderId} [delete]
func (h *StoreHandler) DeleteOrderHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = h.Service.DeleteOrder(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

}
