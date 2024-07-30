package service

import (
	"backend/internal/lib/e"
	"backend/internal/modules/pet/entities"
	"backend/internal/repository"
	"context"
)

type PetService struct {
	DB repository.Repository
}

func NewPetService(db repository.Repository) *PetService {
	return &PetService{DB: db}
}

func (s *PetService) GetById(ctx context.Context, id int) (pet entities.Pet, err error) {
	defer func() { err = e.WrapIfErr("couldn't get pet", err) }()

	if pet, err = s.DB.GetPetById(ctx, id); err != nil {
		return entities.Pet{}, err
	}

	return pet, nil
}

func (s *PetService) UpdateWithForm(id int, name string, status string) error {
	//TODO implement me
	panic("implement me")
}

func (s *PetService) DeleteById(id int) error {
	//TODO implement me
	panic("implement me")
}

func (s *PetService) UploadImage(id int) error {
	//TODO implement me
	panic("implement me")
}

func (s *PetService) Create(pet entities.Pet) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PetService) Update(pet entities.Pet) error {
	//TODO implement me
	panic("implement me")
}

func (s *PetService) GetByStatus(status string) ([]entities.Pet, error) {
	//TODO implement me
	panic("implement me")
}
