package service

import (
	"backend/internal/modules/pet/entities"
	"context"
)

type PetServicer interface {
	GetById(ctx context.Context, id int) (entities.Pet, error)
	UpdateWithForm(ctx context.Context, id int, name string, status string) error
	DeleteById(ctx context.Context, id int) error
	Create(ctx context.Context, pet entities.Pet) (int, error)
	Update(ctx context.Context, pet entities.Pet) error
	GetByStatus(ctx context.Context, status string) ([]entities.Pet, error)
}
