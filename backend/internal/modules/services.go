package modules

import (
	au "backend/internal/modules/auth/service"
	ps "backend/internal/modules/pet/service"
	ss "backend/internal/modules/store/service"
	us "backend/internal/modules/user/service"
	"backend/internal/repository"
)

type Services struct {
	Pet   ps.PetServicer
	User  us.UserServicer
	Store ss.StoreServicer
	Auth  au.AuthServicer
}

func NewServices(db repository.Repository, issuer, audience, secret, cookieDomain string) *Services {
	authService := au.NewAuthService(issuer, audience, secret, cookieDomain)

	return &Services{
		Pet:   ps.NewPetService(db),
		User:  us.NewUserService(db, authService),
		Store: ss.NewStoreService(db),
		Auth:  authService,
	}
}
