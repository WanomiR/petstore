package service

import (
	"backend/internal/mocks/mock_repository"
	"backend/internal/modules/pet/entities"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

type Mocks struct {
	Pets       []entities.Pet
	Categories []entities.Category
	PhotoUrls  []entities.PhotoUrl
	Tags       []entities.Tag
	PetTags    []entities.PetTag
}

func TestPetService_GetById(t *testing.T) {
	testCases := []struct {
		name    string
		petId   int
		wantErr bool
	}{
		{"normal case", 1, false},
		{"unknown pet", 100, true},
		{"unknown category", 4, true},
		{"no photos", 5, false},
		{"no tag pairs", 6, false},
		{"no tag ", 7, true},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := NewMockRepository(controller)
	ps := NewPetService(mockRepo)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := ps.GetById(context.Background(), tc.petId); (err != nil) != tc.wantErr {
				t.Errorf("GetById() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestPetService_UpdateWithForm(t *testing.T) {
	type Payload struct {
		id           int
		name, status string
	}
	testCases := []struct {
		name    string
		payload Payload
		wantErr bool
	}{
		{"normal case", Payload{1, "Bobby", "available"}, false},
		{"unknown pet", Payload{100, "Some Name", ""}, true},
		{"unfilled parameters", Payload{1, "", ""}, true},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := NewMockRepository(controller)
	ps := NewPetService(mockRepo)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := ps.UpdateWithForm(context.Background(), tc.payload.id, tc.payload.name, tc.payload.status); (err != nil) != tc.wantErr {
				t.Errorf("UpdateWithForm() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestPetService_Delete(t *testing.T) {
	testCases := []struct {
		name    string
		petId   int
		wantErr bool
	}{
		{"normal case", 1, false},
		{"unknown pet", 100, true},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := NewMockRepository(controller)
	ps := NewPetService(mockRepo)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := ps.Delete(context.Background(), tc.petId); (err != nil) != tc.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestPetService_Create(t *testing.T) {
	testCases := []struct {
		name    string
		payload entities.Pet
		wantErr bool
	}{
		{
			"normal case",
			entities.Pet{Name: "Pet", Category: entities.Category{Id: 1, Name: "cat"}, Tags: []entities.Tag{{1, "fluffy"}}, Status: "available"},
			false,
		},
		{
			"insufficient data",
			entities.Pet{Name: "", Category: entities.Category{Id: 1, Name: "cat"}, Tags: []entities.Tag{{1, "fluffy"}}, Status: "available"},
			true,
		},
		{
			"invalid status",
			entities.Pet{Name: "Pet", Category: entities.Category{Id: 1, Name: "cat"}, Tags: []entities.Tag{{1, "fluffy"}}, Status: "unknown"},
			true,
		},
		{
			"new category",
			entities.Pet{Name: "Pet", Category: entities.Category{Id: 4, Name: "fish"}, Tags: []entities.Tag{{1, "fluffy"}}, Status: "available"},
			false,
		},
		{
			"new tag",
			entities.Pet{Name: "Pet", Category: entities.Category{Id: 4, Name: "fish"}, Tags: []entities.Tag{{4, "cool"}}, Status: "available"},
			false,
		},
		{
			"with photos",
			entities.Pet{Name: "Pet", Category: entities.Category{Id: 4, Name: "fish"}, PhotoUrls: []string{"https://cool-photo.jpg"}, Status: "available"},
			false,
		},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := NewMockRepository(controller)
	ps := NewPetService(mockRepo)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := ps.Create(context.Background(), tc.payload); (err != nil) != tc.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestPetService_Update(t *testing.T) {
	testCases := []struct {
		name    string
		payload entities.Pet
		wantErr bool
	}{
		{
			"normal case",
			entities.Pet{Id: 1, Name: "Pet", Category: entities.Category{Id: 1, Name: "cat"}, Tags: []entities.Tag{{1, "fluffy"}}, Status: "available"},
			false,
		},
		{
			"unknown pet",
			entities.Pet{Id: 100, Name: "Pet", Category: entities.Category{Id: 1, Name: "cat"}, Tags: []entities.Tag{{1, "fluffy"}}, Status: "available"},
			true,
		},
		{
			"invalid status",
			entities.Pet{Id: 1, Name: "Pet", Category: entities.Category{Id: 1, Name: "cat"}, Tags: []entities.Tag{{1, "fluffy"}}, Status: "unknown"},
			true,
		},
		{
			"new category",
			entities.Pet{Id: 1, Name: "Pet", Category: entities.Category{Id: 4, Name: "fish"}, Tags: []entities.Tag{{1, "fluffy"}}, Status: "available"},
			false,
		},
		{
			"new tag",
			entities.Pet{Id: 1, Name: "Pet", Category: entities.Category{Id: 4, Name: "fish"}, Tags: []entities.Tag{{4, "cool"}}, Status: "available"},
			false,
		},
		{
			"new photos",
			entities.Pet{Id: 1, Name: "Pet", Category: entities.Category{Id: 4, Name: "fish"}, PhotoUrls: []string{"https://cool-photo.jpg", "string"}, Status: "available"},
			false,
		},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := NewMockRepository(controller)
	ps := NewPetService(mockRepo)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := ps.Update(context.Background(), tc.payload); (err != nil) != tc.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestGetPetById(t *testing.T) {
	testCases := []struct {
		name    string
		payload string
		wantErr bool
	}{
		{"normal case", "available", false},
		{"invalid status", "unknown", true},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := NewMockRepository(controller)
	ps := NewPetService(mockRepo)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := ps.GetByStatus(context.Background(), tc.payload); (err != nil) != tc.wantErr {
				t.Errorf("GetByStatus() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func NewMockRepository(controller *gomock.Controller) *mock_repository.MockRepository {
	mockRepo := mock_repository.NewMockRepository(controller)
	mocks := generateMocks()

	mockRepo.EXPECT().GetPetById(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, id int) (entities.Pet, error) {
		for _, pet := range mocks.Pets {
			if pet.Id == id {
				return pet, nil
			}
		}
		return entities.Pet{}, errors.New("pet not found")
	}).AnyTimes()

	mockRepo.EXPECT().GetPetCategoryById(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, id int) (entities.Category, error) {
		for _, cat := range mocks.Categories {
			if cat.Id == id {
				return cat, nil
			}
		}
		return entities.Category{}, errors.New("category not found")
	}).AnyTimes()

	mockRepo.EXPECT().GetPetCategoryByName(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, name string) (entities.Category, error) {
		for _, cat := range mocks.Categories {
			if cat.Name == name {
				return cat, nil
			}
		}
		return entities.Category{}, errors.New("category not found")
	}).AnyTimes()

	mockRepo.EXPECT().GetPhotoUrlsByPetId(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, id int) ([]entities.PhotoUrl, error) {
		result := make([]entities.PhotoUrl, 0)
		for _, photoUrl := range mocks.PhotoUrls {
			if photoUrl.PetId == id {
				result = append(result, photoUrl)
			}
		}
		return result, nil
	}).AnyTimes()

	mockRepo.EXPECT().GetPetTagPairsByPetId(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, id int) ([]entities.PetTag, error) {
		result := make([]entities.PetTag, 0)
		for _, petTag := range mocks.PetTags {
			if petTag.PetId == id {
				result = append(result, petTag)
			}
		}
		return result, nil
	}).AnyTimes()

	mockRepo.EXPECT().GetTagById(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, id int) (entities.Tag, error) {
		for _, tag := range mocks.Tags {
			if tag.Id == id {
				return tag, nil
			}
		}
		return entities.Tag{}, errors.New("tag not found")
	}).AnyTimes()

	mockRepo.EXPECT().GetTagByName(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, name string) (entities.Tag, error) {
		for _, tag := range mocks.Tags {
			if tag.Name == name {
				return tag, nil
			}
		}
		return entities.Tag{}, errors.New("tag not found")
	}).AnyTimes()

	mockRepo.EXPECT().CreateTag(gomock.Any(), gomock.Any()).Return(entities.Tag{Id: 4, Name: "fish"}, nil).AnyTimes()

	mockRepo.EXPECT().UpdatePet(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	mockRepo.EXPECT().DeletePet(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	mockRepo.EXPECT().DeletePetTagsByPetId(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	mockRepo.EXPECT().DeletePhotoUrlsByPetId(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	mockRepo.EXPECT().CreatePetCategory(gomock.Any(), gomock.Any()).Return(entities.Category{}, nil).AnyTimes()

	mockRepo.EXPECT().CreatePet(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(1, nil).AnyTimes()

	mockRepo.EXPECT().CreatePetPhotoUrl(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	mockRepo.EXPECT().CreatePetTagPair(gomock.Any(), gomock.Any(), gomock.Any()).Return(entities.PetTag{}, nil).AnyTimes()

	mockRepo.EXPECT().GetPetsByStatus(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, status string) ([]entities.Pet, error) {
		result := make([]entities.Pet, 0)
		for _, pet := range mocks.Pets {
			if pet.Status == status {
				result = append(result, pet)
			}
		}
		return result, nil
	}).AnyTimes()

	return mockRepo
}

func generateMocks() Mocks {
	return Mocks{
		Pets:       generateMockPets(),
		Categories: generateMockCategories(),
		PhotoUrls:  generatePhotoUrls(),
		Tags:       generateTags(),
		PetTags:    generateMocPetTagPairs(),
	}
}

func generateMockPets() []entities.Pet {
	mockPets := []entities.Pet{
		{
			1,
			entities.Category{Id: 1, Name: "cat"},
			"Poppy",
			[]string{"https://wallbox.ru/wallpapers/main/201152/koshki-392426de15fb.jpg"},
			[]entities.Tag{{1, "fluffy"}},
			"available",
		},
		{
			2,
			entities.Category{Id: 2, Name: "dog"},
			"Abby",
			[]string{"https://wallpapers.com/images/hd/dog-pictures-os09dhwexb80d990.jpg"},
			[]entities.Tag{{2, "calm"}},
			"pending",
		},
		{
			3,
			entities.Category{Id: 3, Name: "rodent"},
			"Basil",
			[]string{"https://i.pinimg.com/originals/59/df/fb/59dffb52f7435ce31979f7e03ce02ab4.jpg"},
			[]entities.Tag{{3, "kind"}},
			"sold",
		},
		{
			4,
			entities.Category{Id: 4, Name: "unknown"},
			"Tilly",
			[]string{"https://i.pinimg.com/originals/59/df/fb/59dffb52f7435ce31979f7e03ce02ab4.jpg"},
			[]entities.Tag{{3, "kind"}},
			"pending",
		},
		{
			5,
			entities.Category{Id: 3, Name: "rodent"},
			"Tilly",
			[]string{""},
			[]entities.Tag{{2, "calm"}},
			"sold",
		},
		{
			6,
			entities.Category{Id: 2, Name: "dog"},
			"Rex",
			[]string{"https://wallbox.ru/wallpapers/main/201152/koshki-392426de15fb.jpg"},
			[]entities.Tag{{3, "calm"}},
			"pending",
		},
		{
			7,
			entities.Category{Id: 1, Name: "cat"},
			"Rex",
			[]string{"https://wallbox.ru/wallpapers/main/201152/koshki-392426de15fb.jpg"},
			[]entities.Tag{{4, "unknown"}},
			"pending",
		},
	}

	return mockPets
}

func generateMockCategories() []entities.Category {
	return []entities.Category{
		{Id: 1, Name: "cat"},
		{Id: 2, Name: "dog"},
		{Id: 3, Name: "rodent"},
	}
}

func generatePhotoUrls() []entities.PhotoUrl {
	return []entities.PhotoUrl{
		{Id: 1, PetId: 1, Url: "https://wallbox.ru/wallpapers/main/201152/koshki-392426de15fb.jpg"},
		{Id: 2, PetId: 2, Url: "https://wallpapers.com/images/hd/dog-pictures-os09dhwexb80d990.jpg"},
		{Id: 3, PetId: 3, Url: "https://i.pinimg.com/originals/59/df/fb/59dffb52f7435ce31979f7e03ce02ab4.jpg"},
		{Id: 4, PetId: 4, Url: "https://i.pinimg.com/originals/59/df/fb/59dffb52f7435ce31979f7e03ce02ab4.jpg"},
		{Id: 5, PetId: 6, Url: "https://wallpapers.com/images/hd/dog-pictures-os09dhwexb80d990.jpg"},
		{Id: 6, PetId: 7, Url: "https://wallpapers.com/images/hd/dog-pictures-os09dhwexb80d990.jpg"},
	}
}

func generateTags() []entities.Tag {
	return []entities.Tag{
		{Id: 1, Name: "fluffy"},
		{Id: 2, Name: "calm"},
		{Id: 3, Name: "kind"},
	}
}

func generateMocPetTagPairs() []entities.PetTag {
	return []entities.PetTag{
		{Id: 1, PetId: 1, TagId: 1},
		{Id: 2, PetId: 2, TagId: 2},
		{Id: 3, PetId: 3, TagId: 3},
		{Id: 4, PetId: 4, TagId: 3},
		{Id: 5, PetId: 5, TagId: 2},
		{Id: 6, PetId: 7, TagId: 4},
	}
}
