package repository

import (
	"backend/internal/modules/pet/entities"
	"context"
	"database/sql"
)

type Repository interface {
	Connection() *sql.DB
	PetRepository
}

type PetRepository interface {
	GetPetById(ctx context.Context, petId int) (entities.Pet, error)
	CreatePet(ctx context.Context, categoryId int, petName string, petStatus string) (int, error)
	UpdatePet(ctx context.Context, pet entities.Pet) error
	DeletePet(ctx context.Context, petId int) error
	GetPetCategoryById(ctx context.Context, categoryId int) (entities.Category, error)
	GetPetCategoryByName(ctx context.Context, categoryName string) (entities.Category, error)
	CreatePetCategory(ctx context.Context, categoryName string) (entities.Category, error)
	GetPhotoUrlsByPetId(ctx context.Context, petId int) ([]entities.PhotoUrl, error)
	DeletePhotoUrlsByPetId(ctx context.Context, petId int) error
	CreatePetPhotoUrl(ctx context.Context, petId int, photoUrl string) error
	GetTagById(ctx context.Context, tagId int) (entities.Tag, error)
	GetTagByName(ctx context.Context, tagName string) (entities.Tag, error)
	CreateTag(ctx context.Context, tagName string) (entities.Tag, error)
	GetPetTagPair(ctx context.Context, petId int, tagId int) (entities.PetTag, error)
	GetPetTagPairsByPetId(ctx context.Context, petId int) ([]entities.PetTag, error)
	DeletePetTagsByPetId(ctx context.Context, petId int) error
	CreatePetTagPair(ctx context.Context, petId int, tagId int) (entities.PetTag, error)
}
