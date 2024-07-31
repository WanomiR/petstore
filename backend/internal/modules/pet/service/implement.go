package service

import (
	"backend/internal/lib/e"
	"backend/internal/modules/pet/entities"
	"backend/internal/repository"
	"context"
	"errors"
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
		return e.Wrap("couldn't update pet", err)
	}

	return nil
}

func (s *PetService) DeleteById(ctx context.Context, id int) error {
	if _, err := s.GetById(ctx, id); err != nil {
		return err
	}

	if err := s.DB.DeletePet(ctx, id); err != nil {
		return e.Wrap("couldn't delete pet", err)
	}

	return nil
}

func (s *PetService) Create(ctx context.Context, pet entities.Pet) (int, error) {
	// check required fields
	if pet.Name == "" || pet.Status == "" || pet.Category.Name == "" {
		return 0, errors.New("pet name or pet status or category name is empty")
	}

	// handle pet category
	category, err := s.DB.GetPetCategoryByName(ctx, pet.Category.Name)
	// use category id if category exists
	if err == nil {
		pet.Category = category
	} else {
		// otherwise create new category
		if category, err = s.DB.CreatePetCategory(ctx, pet.Category.Name); err != nil {
			return 0, e.Wrap("couldn't create pet category", err)
		}
		pet.Category = category
	}

	// create pet and update pet id
	petId, err := s.DB.CreatePet(ctx, pet.Category.Id, pet.Name, pet.Status)
	if err != nil {
		return 0, e.Wrap("couldn't create pet", err)
	}
	pet.Id = petId

	// save photo urls
	for _, url := range pet.PhotoUrls {
		if err = s.DB.CreatePetPhotoUrl(ctx, pet.Id, url); err != nil {
			return 0, e.Wrap("couldn't create pet photo url", err)
		}
	}

	// process tags
	for i := range pet.Tags {
		// check if tag exists
		var tag entities.Tag
		tag, err = s.DB.GetTagByName(ctx, pet.Tags[i].Name)
		if err == nil {
			// use tag id if tag exists
			pet.Tags[i] = tag
		} else {
			// otherwise create new tag
			if tag, err = s.DB.CreateTag(ctx, pet.Tags[i].Name); err != nil {
				return 0, e.Wrap("couldn't create pet tag", err)
			}
			pet.Tags[i] = tag
		}

		// create pet/tag pair
		if _, err = s.DB.CreatePetTagPair(ctx, pet.Id, tag.Id); err != nil {
			return 0, e.Wrap("couldn't create pet tag pair", err)
		}
	}

	return pet.Id, nil
}

func (s *PetService) Update(ctx context.Context, pet entities.Pet) error {
	//TODO implement me
	panic("implement me")
}

func (s *PetService) GetByStatus(ctx context.Context, status string) ([]entities.Pet, error) {
	//TODO implement me
	panic("implement me")
}
