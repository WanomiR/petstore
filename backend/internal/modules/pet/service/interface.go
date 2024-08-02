package service

import (
	"backend/internal/modules/pet/entities"
	"context"
)

//go:generate mockgen -source=./interface.go -destination=../../../mocks/mock_pet_service/mock_pet_service.go

type PetServicer interface {
	GetById(ctx context.Context, id int) (entities.Pet, error)
	UpdateWithForm(ctx context.Context, id int, name string, status string) error
	Delete(ctx context.Context, id int) error
	Create(ctx context.Context, pet entities.Pet) (int, error)
	Update(ctx context.Context, pet entities.Pet) error
	GetByStatus(ctx context.Context, status string) ([]entities.Pet, error)
}
