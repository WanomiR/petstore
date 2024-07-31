package service

import (
	"backend/internal/modules/user/entities"
	"backend/internal/repository"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	DB repository.Repository
	// authentication
}

func NewUserService(db repository.Repository) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) GetUserByName(ctx context.Context, name string) (entities.User, error) {
	user, err := s.DB.GetUserByUsername(ctx, name)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, userUpdate entities.User) error {
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

	// TODO: introduce authentication module
	if userUpdate.Password != "" {
		password, _ := bcrypt.GenerateFromPassword([]byte(userUpdate.Password), bcrypt.DefaultCost)
		user.Password = string(password)
	}

	if err := s.DB.UpdateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) CreateUser(ctx context.Context, user entities.User) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("username and password are mandatory")
	}

	if _, err := s.DB.GetUserByUsername(ctx, user.Username); err == nil {
		return errors.New("user already exists")
	}

	// TODO: introduce authentication module
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)

	if err := s.DB.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil

}

func (s *UserService) DeleteUser(ctx context.Context, username string) error {
	if err := s.DB.DeleteUser(ctx, username); err != nil {
		return err
	}
	return nil
}
