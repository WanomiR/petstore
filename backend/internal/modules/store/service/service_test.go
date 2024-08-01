package service

import (
	"backend/internal/mocks/mock_repository"
	"backend/internal/modules/store/entities"
	"context"
	"github.com/golang/mock/gomock"
	"testing"
)

//go:generate mockgen -source=../../../repository/repository.go -destination=../../../mocks/mock_repository/mock_repository.go

func TestStoreService_CreateOrder(t *testing.T) {
	testCases := []struct {
		name    string
		order   entities.Order
		wantErr bool
	}{
		{"normal case", entities.Order{PetId: 1, Quantity: 1}, false},
		{"invalid pet/order id", entities.Order{PetId: 0, Quantity: 0}, true},
		{"invalid status", entities.Order{PetId: 1, Quantity: 1, Status: "unknown"}, true},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	db := NewMockRepository(controller)
	ss := NewStoreService(db)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := ss.CreateOrder(context.Background(), tc.order); (err != nil) != tc.wantErr {
				t.Errorf("CreateOrder() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestStoreService_GetOrderById(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	db := NewMockRepository(controller)
	ss := NewStoreService(db)

	t.Run("normal case", func(t *testing.T) {
		if _, err := ss.GetOrderById(context.Background(), 1); err != nil {
			t.Errorf("GetOrderById() error = %v, wantErr %v", err, nil)
		}
	})
}

func TestStoreService_DeleteOrder(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	db := NewMockRepository(controller)
	ss := NewStoreService(db)

	t.Run("normal case", func(t *testing.T) {
		if err := ss.DeleteOrder(context.Background(), 1); err != nil {
			t.Errorf("DeleteOrder() error = %v, wantErr %v", err, nil)
		}
	})
}

func NewMockRepository(controller *gomock.Controller) *mock_repository.MockRepository {
	mockDb := mock_repository.NewMockRepository(controller)

	mockDb.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).Return(1, nil).AnyTimes()

	mockDb.EXPECT().GetOrderById(gomock.Any(), gomock.Any()).Return(entities.Order{}, nil).AnyTimes()

	mockDb.EXPECT().DeleteOrder(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	return mockDb
}
