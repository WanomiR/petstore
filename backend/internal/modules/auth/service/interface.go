package service

import (
	"backend/internal/modules/auth/entities"
	"net/http"
)

type AuthServicer interface {
	EncryptPassword(password string) (string, error)
	GenerateTokensPair(subject string) (entities.TokensPair, error)
	GetRefreshCookie(refreshToken string) *http.Cookie
	GetExpiredRefreshCookie() *http.Cookie
	GetTokenFromHeaderAndVerify(w http.ResponseWriter, r *http.Request) (string, *entities.Claims, error)
}
