package service

import (
	"backend/internal/modules/user/entities"
	"backend/internal/repository"
	"context"
)

type UserService struct {
	DB repository.Repository
	// authentication
}

func NewUserService(db repository.Repository) *UserService {
	return &UserService{DB: db}
}

func (u *UserService) GetUserByName(ctx context.Context, name string) (entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) UpdateUser(ctx context.Context, user entities.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) CreateUser(ctx context.Context, user entities.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) DeleteUser(ctx context.Context, username string) error {
	//TODO implement me
	panic("implement me")
}
