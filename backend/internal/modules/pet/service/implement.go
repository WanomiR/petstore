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

func (s *PetService) UpdateWithForm(ctx context.Context, id int, name string, status string) error {
	pet, err := s.GetById(ctx, id)
	if err != nil {
		return err
	}

	if name != "" {
		pet.Name = name
	}
	if status != "" {
		pet.Status = status
	}

	if err = s.DB.UpdatePet(ctx, pet); err != nil {
		return err
	}

	return nil
}

func (s *PetService) DeleteById(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (s *PetService) UploadImage(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (s *PetService) Create(ctx context.Context, pet entities.Pet) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PetService) Update(ctx context.Context, pet entities.Pet) error {
	//TODO implement me
	panic("implement me")
}

func (s *PetService) GetByStatus(ctx context.Context, status string) ([]entities.Pet, error) {
	//TODO implement me
	panic("implement me")
}
