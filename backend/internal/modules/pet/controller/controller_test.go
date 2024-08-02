package controller

import (
	"backend/internal/lib/rr"
	mock_service "backend/internal/mocks/mock_pet_service"
	"backend/internal/modules/pet/entities"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testFilePath string

func init() {
	dir, _ := os.Getwd()
	fileName := "implement.go"
	testFilePath = filepath.Join(dir, fileName)
}

func TestPetControl_GetById(t *testing.T) {
	testCases := []struct {
		name       string
		petId      string
		wantStatus int
	}{
		{"normal case", "1", http.StatusOK},
		{"unknown pet", "10", http.StatusNotFound},
		{"invalid id", "avggjl", http.StatusBadRequest},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockPetService(controller)
	pc := NewPetControl(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/pet/"+tc.petId, nil)
			wr := httptest.NewRecorder()

			pc.GetById(wr, req)

			r := wr.Result()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("Want status %d, got %d", tc.wantStatus, r.StatusCode)
			}

		})
	}
}

func TestPetControl_UpdateWithForm(t *testing.T) {
	type Payload struct{ petId, name, status string }

	testCases := []struct {
		name       string
		payload    Payload
		wantStatus int
	}{
		{"normal case", Payload{"1", "garfield", "available"}, http.StatusOK},
		{"unknown pet", Payload{"10", "garfield", "available"}, http.StatusBadRequest},
		{"invalid id", Payload{"agkla", "garfield", "available"}, http.StatusBadRequest},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockPetService(controller)
	pc := NewPetControl(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			param := url.Values{}
			param.Add("name", tc.payload.name)
			param.Add("status", tc.payload.status)

			req := httptest.NewRequest(http.MethodPost, "/pet/"+tc.payload.petId, strings.NewReader(param.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			wr := httptest.NewRecorder()

			pc.UpdateWithForm(wr, req)

			r := wr.Result()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("Want status %d, got %d", tc.wantStatus, r.StatusCode)
			}

		})
	}
}

func TestPetControl_Delete(t *testing.T) {
	testCases := []struct {
		name       string
		petId      string
		wantStatus int
	}{
		{"normal case", "1", http.StatusOK},
		{"unknown pet", "10", http.StatusNotFound},
		{"invalid id", "avggjl", http.StatusBadRequest},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockPetService(controller)
	pc := NewPetControl(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/pet/"+tc.petId, nil)
			wr := httptest.NewRecorder()

			pc.Delete(wr, req)

			r := wr.Result()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("Want status %d, got %d", tc.wantStatus, r.StatusCode)
			}

		})
	}
}

func TestPetControl_UploadImage(t *testing.T) {
	type Payload struct{ petId, additionalMetadata, filePath string }

	testCases := []struct {
		name       string
		payload    Payload
		wantStatus int
	}{
		{"normal case", Payload{"1", "some metadata", testFilePath}, http.StatusOK},
		{"invalid id", Payload{"agjla", "some metadata", testFilePath}, http.StatusBadRequest},
		{"unknown pet", Payload{"10", "some metadata", testFilePath}, http.StatusNotFound},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockPetService(controller)
	pc := NewPetControl(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			file, _ := os.Open(tc.payload.filePath)
			defer file.Close()

			payload := new(bytes.Buffer)
			writer := multipart.NewWriter(payload)
			_ = writer.WriteField("additionalMetadata", tc.payload.additionalMetadata)

			part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
			_, _ = io.Copy(part, file)
			_ = writer.Close()

			req := httptest.NewRequest(http.MethodPost, "/pet/"+tc.payload.petId+"/uploadImage", payload)
			req.Header.Add("Content-Type", writer.FormDataContentType())
			wr := httptest.NewRecorder()

			pc.UploadImage(wr, req)

			r := wr.Result()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("Want status %d, got %d", tc.wantStatus, r.StatusCode)
			}
		})
	}
}

func TestPetControl_Create(t *testing.T) {
	testCases := []struct {
		name       string
		body       any
		wantStatus int
	}{
		{"normal case", entities.Pet{Name: "garfield", Category: entities.Category{Name: "cat"}, Status: "available"}, http.StatusOK},
		{"invalid payload", entities.Pet{}, http.StatusBadRequest},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockPetService(controller)
	pc := NewPetControl(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			payload := new(bytes.Buffer)
			_ = json.NewEncoder(payload).Encode(tc.body)

			req := httptest.NewRequest(http.MethodPost, "/pet", payload)
			wr := httptest.NewRecorder()

			pc.Create(wr, req)

			r := wr.Result()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("Want status %d, got %d", tc.wantStatus, r.StatusCode)
			}
		})
	}
}

func TestPetControl_Update(t *testing.T) {
	testCases := []struct {
		name       string
		body       any
		wantStatus int
	}{
		{"normal case", entities.Pet{Id: 1, Name: "garfield", Category: entities.Category{Name: "cat"}, Status: "available"}, http.StatusOK},
		{"unknown pet", entities.Pet{Id: 10, Name: "garfield", Category: entities.Category{Name: "cat"}, Status: "available"}, http.StatusNotFound},
		{"insufficient parameters", entities.Pet{Id: 1}, http.StatusBadRequest},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockPetService(controller)
	pc := NewPetControl(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			payload := new(bytes.Buffer)
			_ = json.NewEncoder(payload).Encode(tc.body)

			req := httptest.NewRequest(http.MethodPut, "/pet", payload)
			wr := httptest.NewRecorder()

			pc.Update(wr, req)

			r := wr.Result()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("Want status %d, got %d", tc.wantStatus, r.StatusCode)
			}
		})
	}
}

func TestPetControl_GetByStatus(t *testing.T) {
	testCases := []struct {
		name       string
		petStatus  []string
		wantStatus int
	}{
		{"normal case", []string{"available", "pending"}, http.StatusOK},
		{"invalid status", []string{"unknown"}, http.StatusBadRequest},
		{"no status at all", []string{""}, http.StatusBadRequest},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockPetService(controller)
	pc := NewPetControl(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/pet/findByStatus?status="+strings.Join(tc.petStatus, ","), nil)
			wr := httptest.NewRecorder()

			pc.GetByStatus(wr, req)

			r := wr.Result()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("Want status %d, got %d", tc.wantStatus, r.StatusCode)
			}

		})
	}
}

func NewMockPetService(controller *gomock.Controller) *mock_service.MockPetServicer {
	mockService := mock_service.NewMockPetServicer(controller)

	mockService.EXPECT().GetById(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, petId int) (entities.Pet, error) {
		if petId < 1 || petId > 9 {
			return entities.Pet{}, errors.New("pet not found")
		}
		return entities.Pet{}, nil
	}).AnyTimes()

	mockService.EXPECT().UpdateWithForm(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, id int, name, status string) error {
		if id < 1 || id > 9 {
			return errors.New("pet not found")
		}
		return nil
	}).AnyTimes()

	mockService.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	mockService.EXPECT().Create(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, pet entities.Pet) (int, error) {
		if pet.Name == "" || pet.Status == "" || pet.Category.Name == "" {
			return 0, errors.New("fill in mandatory fields")
		}
		return 1, nil
	}).AnyTimes()

	mockService.EXPECT().Update(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, pet entities.Pet) error {
		if pet.Name == "" || pet.Status == "" || pet.Category.Name == "" {
			return errors.New("fill in mandatory fields")
		}
		return nil
	}).AnyTimes()

	mockService.EXPECT().GetByStatus(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, status string) ([]entities.Pet, error) {
		if !(status == "available" || status == "pending" || status == "sold") {
			return nil, errors.New("invalid status")
		}
		return []entities.Pet{{Name: "garfield", Status: "available"}}, nil
	}).AnyTimes()

	return mockService
}
