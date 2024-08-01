package service

import (
	ue "backend/internal/modules/user/entities"
	"context"
	"net/http"
)

type UserServicer interface {
	GetByName(ctx context.Context, name string) (ue.User, error)
	Update(ctx context.Context, user ue.User) error
	Create(ctx context.Context, user ue.User) error
	Delete(ctx context.Context, username string) error
	Authorize(ctx context.Context, username, password string) (string, *http.Cookie, error)
	ResetCookie() *http.Cookie
}
