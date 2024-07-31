package service

import (
	"backend/internal/modules/user/entities"
	"context"
)

type UserServicer interface {
	GetUserByName(ctx context.Context, name string) (entities.User, error)
	UpdateUser(ctx context.Context, user entities.User) error
	CreateUser(ctx context.Context, user entities.User) error
	DeleteUser(ctx context.Context, username string) error
}
