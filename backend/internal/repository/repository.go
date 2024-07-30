package repository

import (
	"backend/internal/modules/pet/entities"
	"context"
	"database/sql"
)

type Repository interface {
	Connection() *sql.DB
	GetPetById(ctx context.Context, petId int) (entities.Pet, error)
}
