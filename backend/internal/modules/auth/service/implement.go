package service

import (
	"backend/internal/modules/auth/entities"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type AuthService struct {
	Issuer        string // host
	Audience      string // host
	Secret        string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string // host
	CookiePath    string
	CookieName    string
}

func NewAuthService(issuer, audience, secret, cookieDomain string) *AuthService {
	return &AuthService{
		Issuer:        issuer,
		Audience:      audience,
		Secret:        secret,
		TokenExpiry:   15 * time.Minute,
		RefreshExpiry: 24 * time.Hour,
		CookieDomain:  cookieDomain,
		CookiePath:    "/",
		CookieName:    "__Host-refresh_token",
	}
}

func (a *AuthService) EncryptPassword(password string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encrypted), nil
}

func (a *AuthService) GenerateTokensPair(subject string) (entities.TokensPair, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthService) GetRefreshCookie(refreshToken string) *http.Cookie {
	//TODO implement me
	panic("implement me")
}

func (a *AuthService) GetExpiredRefreshCookie() *http.Cookie {
	//TODO implement me
	panic("implement me")
}

func (a *AuthService) GetTokenFromHeaderAndVerify(w http.ResponseWriter, r *http.Request) (string, *entities.Claims, error) {
	//TODO implement me
	panic("implement me")
}
