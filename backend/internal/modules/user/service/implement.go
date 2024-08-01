package service

import (
	"backend/internal/lib/e"
	"backend/internal/modules/auth/service"
	"backend/internal/modules/user/entities"
	"backend/internal/repository"
	"context"
	"errors"
	"net/http"
)

type UserService struct {
	DB   repository.Repository
	auth service.AuthServicer
}

func NewUserService(db repository.Repository, auth service.AuthServicer) *UserService {
	return &UserService{DB: db, auth: auth}
}

func (s *UserService) GetByName(ctx context.Context, name string) (entities.User, error) {
	user, err := s.DB.GetUserByUsername(ctx, name)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (s *UserService) Update(ctx context.Context, userUpdate entities.User) error {
	user, _ := s.DB.GetUserByUsername(ctx, userUpdate.Username)

	if userUpdate.FirstName != "" {
		user.FirstName = userUpdate.FirstName
	}

	if userUpdate.LastName != "" {
		user.LastName = userUpdate.LastName
	}

	if userUpdate.Email != "" {
		user.Email = userUpdate.Email
	}

	if userUpdate.Phone != "" {
		user.Phone = userUpdate.Phone
	}

	if userUpdate.UserStatus != user.UserStatus {
		user.UserStatus = userUpdate.UserStatus
	}

	if userUpdate.Password != "" {
		user.Password, _ = s.auth.EncryptPassword(userUpdate.Password)
	}

	if err := s.DB.UpdateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) Create(ctx context.Context, user entities.User) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("username and password are mandatory")
	}

	if _, err := s.DB.GetUserByUsername(ctx, user.Username); err == nil {
		return errors.New("user already exists")
	}

	user.Password, _ = s.auth.EncryptPassword(user.Password)

	if err := s.DB.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil

}

func (s *UserService) Delete(ctx context.Context, username string) error {
	if err := s.DB.DeleteUser(ctx, username); err != nil {
		return err
	}
	return nil
}

func (s *UserService) Authorize(ctx context.Context, username, password string) (string, *http.Cookie, error) {
	user, err := s.GetByName(ctx, username)
	if err != nil {
		return "", nil, e.Wrap("user not found", err)
	}

	ok, err := s.auth.VerifyPassword(password, user.Password)
	if err != nil || !ok {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := s.auth.GenerateToken(username)
	if err != nil {
		return "", nil, e.Wrap("failed to generate tokens", err)
	}

	refreshCookie := s.auth.CreateCookie(token)

	return token, refreshCookie, nil
}

func (s *UserService) ResetCookie() *http.Cookie {
	return s.auth.CreateExpiredCookie()
}
