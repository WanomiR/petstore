package modules

import (
	au "backend/internal/modules/auth/service"
	ps "backend/internal/modules/pet/service"
	us "backend/internal/modules/user/service"
	"backend/internal/repository"
)

type Services struct {
	Pet  ps.PetServicer
	User us.UserServicer
	Auth au.AuthServicer
}

func NewServices(db repository.Repository, issuer, audience, secret, cookieDomain string) *Services {
	return &Services{
		Pet:  ps.NewPetService(db),
		User: us.NewUserService(db),
		Auth: au.NewAuthService(issuer, audience, secret, cookieDomain),
	}
}
