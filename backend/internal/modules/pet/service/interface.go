package service

import "backend/internal/modules/pet/entities"

type PetServicer interface {
	GetById(id int) (entities.Pet, error)
	UpdateWithForm(id int, name string, status string) error
	DeleteById(id int) error
	UploadImage(id int) error
	Create(pet entities.Pet) (int, error)
	Update(pet entities.Pet) error
	GetByStatus(status string) ([]entities.Pet, error)
}
