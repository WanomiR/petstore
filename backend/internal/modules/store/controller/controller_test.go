package controller

import (
	"backend/internal/lib/rr"
	"backend/internal/mocks/mock_store_service"
	"backend/internal/modules/store/entities"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:generate mockgen -source=../service/interface.go -destination=../../../mocks/mock_store_service/mock_store_service.go

func TestStoreControl_GetInventory(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockService(controller)
	sc := NewStoreControl(mockService, rr.NewReadRespond())

	req := httptest.NewRequest(http.MethodGet, "/store/inventory", nil)
	wr := httptest.NewRecorder()

	sc.GetInventory(wr, req)

	r := wr.Result()

	if r.StatusCode != http.StatusOK {
		t.Run("normal case", func(t *testing.T) {
			t.Errorf("want status code %d, got %d", http.StatusOK, r.StatusCode)
		})
	}
}

func TestStoreControl_CreateOrder(t *testing.T) {
	testCases := []struct {
		name       string
		body       any
		wantStatus int
	}{
		{"normal case", entities.Order{PetId: 1, Quantity: 2}, http.StatusOK},
		{"invalid case", entities.Order{PetId: 0, Quantity: 0}, http.StatusBadRequest},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockService(controller)
	sc := NewStoreControl(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var body bytes.Buffer
			json.NewEncoder(&body).Encode(tc.body)

			req := httptest.NewRequest(http.MethodPost, "/store/order", &body)
			wr := httptest.NewRecorder()

			sc.CreateOrder(wr, req)

			r := wr.Result()

			var resp rr.JSONResponse
			json.NewDecoder(r.Body).Decode(&resp)
			defer r.Body.Close()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("CreateOrder(), expected status %d, got %d", tc.wantStatus, r.StatusCode)
			}
		})
	}
}

func TestStoreControl_GetOrderById(t *testing.T) {
	testCases := []struct {
		name       string
		orderId    string
		wantStatus int
	}{
		{"normal case", "1", 200},
		{"order id out of range", "20", 404},
		{"invalid order id", "klajg", 400},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockService(controller)
	sc := NewStoreControl(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/store/order/"+tc.orderId, nil)
			wr := httptest.NewRecorder()

			sc.GetOrderById(wr, req)

			r := wr.Result()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("GetOrderById(), expected status %d, got %d", tc.wantStatus, r.StatusCode)
			}
		})
	}
}

func TestStoreControl_DeleteOrder(t *testing.T) {
	testCases := []struct {
		name       string
		orderId    string
		wantStatus int
	}{
		{"normal case", "1", 200},
		{"order id out of range", "20", 404},
		{"invalid order id", "klajg", 400},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockService(controller)
	sc := NewStoreControl(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/store/order/"+tc.orderId, nil)
			wr := httptest.NewRecorder()

			sc.DeleteOrder(wr, req)

			r := wr.Result()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("DeleteOrder(), expected status %d, got %d", tc.wantStatus, r.StatusCode)
			}
		})
	}
}

func NewMockService(controller *gomock.Controller) *mock_service.MockStoreServicer {
	mockService := mock_service.NewMockStoreServicer(controller)

	mockService.EXPECT().GetOrderById(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, id int) (entities.Order, error) {
		if id < 1 || id > 9 {
			return entities.Order{}, errors.New("order not found")
		}
		return entities.Order{}, nil
	}).AnyTimes()

	mockService.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, order entities.Order) (int, error) {
		if order.PetId == 0 || order.Quantity == 0 {
			return 0, errors.New("invalid order")
		}
		return 1, nil
	}).AnyTimes()

	mockService.EXPECT().DeleteOrder(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	return mockService
}
