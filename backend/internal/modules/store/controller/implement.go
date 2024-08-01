package controller

import (
	"backend/internal/lib/e"
	"backend/internal/lib/rr"
	"backend/internal/lib/u"
	"backend/internal/modules/store/entities"
	"backend/internal/modules/store/service"
	"fmt"
	"net/http"
	"strconv"
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

// CreateOrder godoc
// @Summary create order
// @Description Place an order for a pet
// @Tags store
// @Accept json
// @Produce json
// @Param body body entities.Order true "Order placed for purchasing a pet"
// @Success 200 {object} rr.JSONResponse
// @Failure 400 {object} rr.JSONResponse
// @Router /store/order [post]
func (s *StoreControl) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order entities.Order
	_ = s.rr.ReadJSON(w, r, &order)

	orderId, err := s.service.CreateOrder(r.Context(), order)
	if err != nil {
		_ = s.rr.WriteJSONError(w, e.Wrap("couldn't create order", err))
		return
	}

	resp := rr.JSONResponse{Error: false, Message: fmt.Sprintf("order created, id: %d", orderId)}
	_ = s.rr.WriteJSON(w, 200, resp)
}

// GetOrderById godoc
// @Summary get order
// @Description Fund purchase order by id
// @Tags store
// @Produce json
// @Param orderId path int true "ID of order that needs to be fetched"
// @Success 200 {object} entities.Order
// @Failure 400,404 {object} rr.JSONResponse
// @Router /store/order/{orderId} [get]
func (s *StoreControl) GetOrderById(w http.ResponseWriter, r *http.Request) {
	id := u.ParamFromPath(r.URL.Path)

	orderId, err := strconv.Atoi(id)
	if err != nil {
		_ = s.rr.WriteJSONError(w, e.Wrap("invalid id supplied", err), 400)
		return
	}

	order, err := s.service.GetOrderById(r.Context(), orderId)
	if err != nil {
		_ = s.rr.WriteJSONError(w, e.Wrap("couldn't get order", err), 404)
		return
	}

	_ = s.rr.WriteJSON(w, 200, order)
}

// DeleteOrder godoc
// @Summary delete order
// @Description Delete purchase order by id
// @Tags store
// @Produce json
// @Param orderId path int true "ID of the order that needs to be deleted"
// @Success 200 {object} rr.JSONResponse
// @Failure 400,404 {object} rr.JSONResponse
// @Router /store/order/{orderId} [delete]
func (s *StoreControl) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := u.ParamFromPath(r.URL.Path)

	orderId, err := strconv.Atoi(id)
	if err != nil {
		_ = s.rr.WriteJSONError(w, e.Wrap("invalid id supplied", err))
		return
	}

	if _, err = s.service.GetOrderById(r.Context(), orderId); err != nil {
		_ = s.rr.WriteJSONError(w, e.Wrap("couldn't get order", err), 404)
		return
	}

	if err = s.service.DeleteOrder(r.Context(), orderId); err != nil {
		_ = s.rr.WriteJSONError(w, e.Wrap("couldn't delete order", err))
		return
	}

	resp := rr.JSONResponse{Error: false, Message: "order deleted"}
	_ = s.rr.WriteJSON(w, 200, resp)
}
