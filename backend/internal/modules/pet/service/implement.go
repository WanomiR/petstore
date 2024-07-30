package service

import (
	"backend/internal/modules/pet/entities"
	"backend/internal/repository"
)

type PetService struct {
	DB repository.Repository
}

func NewPetService(db repository.Repository) *PetService {
	return &PetService{DB: db}
}

func (p *PetService) GetById(id int) (entities.Pet, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PetService) UpdateWithForm(id int, name string, status string) error {
	//TODO implement me
	panic("implement me")
}

func (p *PetService) DeleteById(id int) error {
	//TODO implement me
	panic("implement me")
}

func (p *PetService) UploadImage(id int) error {
	//TODO implement me
	panic("implement me")
}

func (p *PetService) Create(pet entities.Pet) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PetService) Update(pet entities.Pet) error {
	//TODO implement me
	panic("implement me")
}

func (p *PetService) GetByStatus(status string) ([]entities.Pet, error) {
	//TODO implement me
	panic("implement me")
}
