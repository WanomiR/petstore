package controller

import (
	"backend/internal/lib/rr"
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

func (s *StoreControl) GetInventory(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
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
