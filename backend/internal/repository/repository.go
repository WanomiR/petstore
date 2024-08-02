package repository

import (
	pe "backend/internal/modules/pet/entities"
	se "backend/internal/modules/store/entities"
	ue "backend/internal/modules/user/entities"
	"context"
	"database/sql"
)

//go:generate mockgen -source=./repository.go -destination=../mocks/mock_repository/mock_repository.go

type Repository interface {
	Connection() *sql.DB
	UserRepository
	PetRepository
	StoreRepository
}
type PetRepository interface {
	GetPetById(ctx context.Context, petId int) (pe.Pet, error)
	CreatePet(ctx context.Context, categoryId int, petName string, petStatus string) (int, error)
	UpdatePet(ctx context.Context, pet pe.Pet) error
	DeletePet(ctx context.Context, petId int) error
	GetPetsByStatus(ctx context.Context, petStatus string) ([]pe.Pet, error)
	GetPetCategoryById(ctx context.Context, categoryId int) (pe.Category, error)
	GetPetCategoryByName(ctx context.Context, categoryName string) (pe.Category, error)
	CreatePetCategory(ctx context.Context, categoryName string) (pe.Category, error)
	GetPhotoUrlsByPetId(ctx context.Context, petId int) ([]pe.PhotoUrl, error)
	DeletePhotoUrlsByPetId(ctx context.Context, petId int) error
	CreatePetPhotoUrl(ctx context.Context, petId int, photoUrl string) error
	GetTagById(ctx context.Context, tagId int) (pe.Tag, error)
	GetTagByName(ctx context.Context, tagName string) (pe.Tag, error)
	CreateTag(ctx context.Context, tagName string) (pe.Tag, error)
	GetPetTagPair(ctx context.Context, petId int, tagId int) (pe.PetTag, error)
	GetPetTagPairsByPetId(ctx context.Context, petId int) ([]pe.PetTag, error)
	DeletePetTagsByPetId(ctx context.Context, petId int) error
	CreatePetTagPair(ctx context.Context, petId int, tagId int) (pe.PetTag, error)
}

type StoreRepository interface {
	GetOrderById(ctx context.Context, orderId int) (se.Order, error)
	CreateOrder(ctx context.Context, order se.Order) (int, error)
	DeleteOrder(ctx context.Context, orderId int) error
}

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (ue.User, error)
	UpdateUser(ctx context.Context, user ue.User) error
	CreateUser(ctx context.Context, user ue.User) (int, error)
	DeleteUser(ctx context.Context, username string) error
}
