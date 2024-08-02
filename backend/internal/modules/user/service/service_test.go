package service

import (
	mock_service "backend/internal/mocks/mock_auth_service"
	"backend/internal/mocks/mock_repository"
	"backend/internal/modules/user/entities"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
)

func TestUserService_GetByName(t *testing.T) {
	testCases := []struct {
		name     string
		username string
		wantErr  bool
	}{
		{"normal case", "wanomir", false},
		{"unknown user", "john", true},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := NewMockRepository(controller)
	mockAuth := NewMockAuth(controller)

	us := NewUserService(mockRepo, mockAuth)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := us.GetByName(context.Background(), tc.username); (err != nil) != tc.wantErr {
				t.Errorf("GetByName() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestUserService_Update(t *testing.T) {
	testCases := []struct {
		name       string
		userUpdate entities.User
		wantErr    bool
	}{
		{"normal case", entities.User{FirstName: "John", LastName: "Snow", Email: "j.snow@wfell.com", Phone: "7-999-412-42-42", Password: "password", UserStatus: 1}, false},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := NewMockRepository(controller)
	mockAuth := NewMockAuth(controller)

	us := NewUserService(mockRepo, mockAuth)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := us.Update(context.Background(), tc.userUpdate); (err != nil) != tc.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestUserService_Create(t *testing.T) {
	testCases := []struct {
		name    string
		user    entities.User
		wantErr bool
	}{
		{"normal case", entities.User{Username: "john", Password: "password"}, false},
		{"incomplete credentials", entities.User{Username: "", Password: ""}, true},
		{"user exists", entities.User{Username: "wanomir", Password: "password"}, true},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := NewMockRepository(controller)
	mockAuth := NewMockAuth(controller)

	us := NewUserService(mockRepo, mockAuth)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := us.Create(context.Background(), tc.user); (err != nil) != tc.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestUserService_Delete(t *testing.T) {
	testCases := []struct {
		name     string
		username string
		wantErr  bool
	}{
		{"normal case", "wanomir", false},
		{"unknown user", "john", true},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := NewMockRepository(controller)
	mockAuth := NewMockAuth(controller)

	us := NewUserService(mockRepo, mockAuth)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := us.Delete(context.Background(), tc.username); (err != nil) != tc.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestUserService_Authorize(t *testing.T) {
	testCases := []struct {
		name     string
		username string
		password string
		wantErr  bool
	}{
		{"normal case", "wanomir", "password", false},
		{"unknown user", "john", "password", true},
		{"password mismatch", "wanomir", "my-password", true},
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockRepo := NewMockRepository(controller)
	mockAuth := NewMockAuth(controller)

	us := NewUserService(mockRepo, mockAuth)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if _, _, err := us.Authorize(context.Background(), tc.username, tc.password); (err != nil) != tc.wantErr {
				t.Errorf("Authorize() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func NewMockAuth(controller *gomock.Controller) *mock_service.MockAuthServicer {
	mockAuth := mock_service.NewMockAuthServicer(controller)

	mockAuth.EXPECT().EncryptPassword(gomock.Any()).Return("password", nil).AnyTimes()

	mockAuth.EXPECT().VerifyPassword(gomock.Any(), gomock.Any()).DoAndReturn(func(password, _ string) (bool, error) {
		if password != "password" {
			return false, nil
		}
		return true, nil
	}).AnyTimes()

	mockAuth.EXPECT().GenerateToken(gomock.Any()).Return("token", nil).AnyTimes()

	mockAuth.EXPECT().CreateCookie(gomock.Any()).Return(&http.Cookie{}).AnyTimes()

	mockAuth.EXPECT().CreateExpiredCookie().Return(&http.Cookie{}).AnyTimes()

	return mockAuth
}

func NewMockRepository(controller *gomock.Controller) *mock_repository.MockRepository {
	mockDb := mock_repository.NewMockRepository(controller)

	mockDb.EXPECT().GetUserByUsername(gomock.Any(), gomock.Any()).DoAndReturn(func(_ctx context.Context, name string) (entities.User, error) {
		if name == "wanomir" || name == "jenstar" {
			return entities.User{}, nil
		}
		return entities.User{}, errors.New("user not found")
	}).AnyTimes()

	mockDb.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	mockDb.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(1, nil).AnyTimes()

	mockDb.EXPECT().DeleteUser(gomock.Any(), gomock.Any()).DoAndReturn(func(_ctx context.Context, name string) error {
		if name == "wanomir" || name == "jenstar" {
			return nil
		}
		return errors.New("user not found")
	}).AnyTimes()

	return mockDb
}
