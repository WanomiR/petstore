package service

import (
	"backend/internal/lib/e"
	"backend/internal/modules/pet/entities"
	"backend/internal/repository"
	"context"
	"errors"
	"log"
)

type PetService struct {
	DB repository.Repository
}

func NewPetService(db repository.Repository) *PetService {
	return &PetService{DB: db}
}

func (s *PetService) GetById(ctx context.Context, id int) (pet entities.Pet, err error) {
	if pet, err = s.DB.GetPetById(ctx, id); err != nil {
		return entities.Pet{}, err
	}

	// get pet category
	if pet.Category, err = s.DB.GetPetCategoryById(ctx, pet.Category.Id); err != nil {
		return pet, e.Wrap("failed to get pet category", err)
	}

	// get pet photo urls
	photoUrls, err := s.DB.GetPhotoUrlsByPetId(ctx, pet.Id)
	if err != nil {
		return pet, e.Wrap("failed to get photo urls", err)
	}

	pet.PhotoUrls = make([]string, 0, len(photoUrls))
	for _, photoUrl := range photoUrls {
		pet.PhotoUrls = append(pet.PhotoUrls, photoUrl.Url)
	}

	// get pet tag pairs
	petTagPairs, err := s.DB.GetPetTagPairsByPetId(ctx, pet.Id)
	if err != nil {
		return entities.Pet{}, e.Wrap("failed to get pet tag pairs", err)
	}

	// get tags for current pet
	pet.Tags = make([]entities.Tag, 0)
	for _, pt := range petTagPairs {
		var tag entities.Tag
		if tag, err = s.DB.GetTagById(ctx, pt.TagId); err != nil {
			log.Println("failed to get tag", err.Error())
			continue
		}
		pet.Tags = append(pet.Tags, tag)
	}

	return pet, nil
}

func (s *PetService) UpdateWithForm(ctx context.Context, id int, name string, status string) error {
	pet, err := s.GetById(ctx, id)
	if err != nil {
		return e.Wrap("couldn't get pet", err)
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
		return e.Wrap("couldn't get pet", err)
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
			// otherwise create a new tag
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

func (s *PetService) Update(ctx context.Context, petUpdate entities.Pet) error {
	pet, err := s.GetById(ctx, petUpdate.Id)
	if err != nil {
		return e.Wrap("couldn't get pet", err)
	}

	if petUpdate.Name != "" {
		pet.Name = petUpdate.Name
	}

	if petUpdate.Status != "" {
		pet.Status = petUpdate.Status
	}

	if petUpdate.Category.Name != "" {
		var category entities.Category

		category, err = s.DB.GetPetCategoryByName(ctx, petUpdate.Category.Name)
		// use category id if category exists
		if err == nil {
			pet.Category = category
		} else {
			// otherwise create new category
			if category, err = s.DB.CreatePetCategory(ctx, pet.Category.Name); err != nil {
				return e.Wrap("couldn't create pet category", err)
			}
			pet.Category = category
		}
	}

	// update pet
	if err = s.DB.UpdatePet(ctx, pet); err != nil {
		return e.Wrap("couldn't update pet", err)
	}

	pet.PhotoUrls = petUpdate.PhotoUrls
	pet.Tags = petUpdate.Tags

	// update photo urls
	if len(pet.PhotoUrls) > 0 {
		// delete all current photo urls
		if err = s.DB.DeletePhotoUrlsByPetId(ctx, pet.Id); err != nil {
			return e.Wrap("couldn't delete photo urls", err)
		}
		// iterate over new photo urls
		for _, url := range pet.PhotoUrls {
			if err = s.DB.CreatePetPhotoUrl(ctx, pet.Id, url); err != nil {
				return e.Wrap("couldn't create pet photo url", err)
			}
		}
	}

	// update tags
	if len(pet.Tags) > 0 {
		// delete existing tags for this pet
		if err = s.DB.DeletePetTagsByPetId(ctx, pet.Id); err != nil {
			return e.Wrap("couldn't delete pet tags", err)
		}
		// iterate over new tags
		for i := range pet.Tags {
			// check if tag exists
			var tag entities.Tag
			tag, err = s.DB.GetTagByName(ctx, pet.Tags[i].Name)
			if err == nil {
				// use tag id if tag exists
				pet.Tags[i] = tag
			} else {
				// otherwise create a new tag
				if tag, err = s.DB.CreateTag(ctx, pet.Tags[i].Name); err != nil {
					return e.Wrap("couldn't create pet tag", err)
				}
				pet.Tags[i] = tag
			}
			// create pet/tag pair
			if _, err = s.DB.CreatePetTagPair(ctx, pet.Id, tag.Id); err != nil {
				return e.Wrap("couldn't create pet tag pair", err)
			}
		}
	}

	return nil
}

func (s *PetService) GetByStatus(ctx context.Context, status string) ([]entities.Pet, error) {
	//TODO implement me
	panic("implement me")
}
