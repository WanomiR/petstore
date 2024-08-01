package controller

import (
	"backend/internal/lib/rr"
	"backend/internal/modules/store/entities"
	"backend/internal/modules/store/service"
	"net/http"
)

type StoreControl struct {
	service service.StoreServicer
	rr      rr.ReadResponder
}

func NewStoreControl(service service.StoreServicer, readResponder rr.ReadResponder) *StoreControl {
	return &StoreControl{service: service, rr: readResponder}
}

// GetInventory godoc
// @Summary get inventory
// @Security ApiKeyAuth
// @Description Returns pet inventories
// @Tags store
// @Produce json
// @Success 200 {object} entities.Inventory
// @Router /store/inventory [get]
func (s *StoreControl) GetInventory(w http.ResponseWriter, r *http.Request) {
	payload := entities.Inventory{
		"additionalProp1": 0,
		"additionalProp2": 0,
		"additionalProp3": 0,
	}
	_ = s.rr.WriteJSON(w, 200, payload)
}

func (s *StoreControl) CreateOrder(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s *StoreControl) GetOrderById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s *StoreControl) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
