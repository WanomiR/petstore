package controller

import (
	"backend/internal/lib/rr"
	mock_service "backend/internal/mocks/mock_user_service"
	"backend/internal/modules/user/entities"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserControl_GetByUsername(t *testing.T) {
	testCases := []struct {
		name       string
		username   string
		wantStatus int
	}{
		{"normal case", "wanomir", http.StatusOK},
		{"invalid name", "", http.StatusBadRequest},
		{"unknown user", "john", http.StatusNotFound},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockUserService(controller)
	uc := NewUserController(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/user/"+tc.username, nil)
			wr := httptest.NewRecorder()

			uc.GetByUsername(wr, req)
			r := wr.Result()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("want status %d, got %d", tc.wantStatus, r.StatusCode)
			}
		})
	}
}

func TestUserControl_Update(t *testing.T) {
	testCases := []struct {
		name       string
		username   string
		body       any
		wantStatus int
	}{
		{"normal case", "wanomir", entities.User{Username: "wanomir"}, http.StatusOK},
		{"invalid username", "", entities.User{Username: "wanomir"}, http.StatusBadRequest},
		{"unknown user", "john", entities.User{Username: "john"}, http.StatusNotFound},
		{"username mismatch", "wanomir", entities.User{Username: "john"}, http.StatusBadRequest},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockUserService(controller)
	uc := NewUserController(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var body bytes.Buffer
			json.NewEncoder(&body).Encode(tc.body)

			req := httptest.NewRequest(http.MethodPut, "/user/"+tc.username, &body)
			wr := httptest.NewRecorder()

			uc.Update(wr, req)
			r := wr.Result()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("want status %d, got %d", tc.wantStatus, r.StatusCode)
			}

		})
	}
}

func TestUserControl_Create(t *testing.T) {
	testCases := []struct {
		name       string
		body       any
		wantStatus int
	}{
		{"normal case", entities.User{Username: "john", Password: "snow"}, http.StatusCreated},
		{"missing password", entities.User{Username: "john", Password: ""}, http.StatusBadRequest},
		{"user already exists", entities.User{Username: "wanomir", Password: "password"}, http.StatusBadRequest},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockUserService(controller)
	uc := NewUserController(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var body bytes.Buffer
			json.NewEncoder(&body).Encode(tc.body)

			req := httptest.NewRequest(http.MethodPost, "/user", &body)
			wr := httptest.NewRecorder()

			uc.Create(wr, req)
			r := wr.Result()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("want status %d, got %d", tc.wantStatus, r.StatusCode)
			}

		})
	}
}

func TestUserControl_Delete(t *testing.T) {
	testCases := []struct {
		name       string
		username   string
		wantStatus int
	}{
		{"normal case", "wanomir", http.StatusOK},
		{"invalid name", "", http.StatusBadRequest},
		{"unknown user", "john", http.StatusNotFound},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockUserService(controller)
	uc := NewUserController(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/user/"+tc.username, nil)
			wr := httptest.NewRecorder()

			uc.Delete(wr, req)
			r := wr.Result()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("want status %d, got %d", tc.wantStatus, r.StatusCode)
			}
		})
	}
}

func TestUserControl_CreateWithArray(t *testing.T) {
	testCases := []struct {
		name       string
		body       any
		wantStatus int
	}{
		{"normal case", entities.Users{entities.User{Username: "john", Password: "snow"}}, http.StatusCreated},
		{"missing password", entities.Users{entities.User{Username: "john", Password: ""}}, http.StatusBadRequest},
		{"user already exists", entities.Users{entities.User{Username: "wanomir", Password: "password"}}, http.StatusBadRequest},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockUserService(controller)
	uc := NewUserController(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var body bytes.Buffer
			json.NewEncoder(&body).Encode(tc.body)

			req := httptest.NewRequest(http.MethodPost, "/user/createWithArray", &body)
			wr := httptest.NewRecorder()

			uc.CreateWithArray(wr, req)
			r := wr.Result()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("want status %d, got %d", tc.wantStatus, r.StatusCode)
			}

		})
	}
}

func TestUserControl_Login(t *testing.T) {
	testCases := []struct {
		name       string
		username   string
		password   string
		wantStatus int
	}{
		{"normal case", "wanomir", "password", http.StatusOK},
		{"incorrect credentials", "jenstar", "my-password", http.StatusUnauthorized},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockUserService(controller)
	uc := NewUserController(mockService, rr.NewReadRespond())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodGet, "/user/login?username="+tc.username+"&password="+tc.password, nil)
			wr := httptest.NewRecorder()

			uc.Login(wr, req)
			r := wr.Result()

			if r.StatusCode != tc.wantStatus {
				t.Errorf("want status %d, got %d", tc.wantStatus, r.StatusCode)
			}

		})
	}
}

func TestUserControl_Logout(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := NewMockUserService(controller)
	uc := NewUserController(mockService, rr.NewReadRespond())

	t.Run("normal case", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/user/logout", nil)
		wr := httptest.NewRecorder()

		uc.Logout(wr, req)
		r := wr.Result()

		if r.StatusCode != http.StatusOK {
			t.Errorf("want status %d, got %d", http.StatusOK, r.StatusCode)
		}
	})
}

func NewMockUserService(controller *gomock.Controller) *mock_service.MockUserServicer {
	mockService := mock_service.NewMockUserServicer(controller)

	mockService.EXPECT().GetByName(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, name string) (entities.User, error) {
		if name == "wanomir" || name == "jenstar" {
			return entities.User{}, nil
		}
		return entities.User{}, errors.New("user not found")
	}).AnyTimes()

	mockService.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	mockService.EXPECT().Create(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, user entities.User) (int, error) {
		if user.Username == "" || user.Password == "" {
			return 0, errors.New("need username and password")
		}

		if user.Username == "wanomir" || user.Username == "jenstar" {
			return 0, errors.New("user already exists")
		}

		return 1, nil
	}).AnyTimes()

	mockService.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	mockService.EXPECT().Authorize(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, username, password string) (string, *http.Cookie, error) {
		if (username == "wanomir" || username == "jenstar") && password == "password" {
			return "token", &http.Cookie{}, nil
		}
		return "", &http.Cookie{}, errors.New("unauthorized user")
	}).AnyTimes()

	mockService.EXPECT().ResetCookie().Return(&http.Cookie{}).AnyTimes()

	return mockService
}
