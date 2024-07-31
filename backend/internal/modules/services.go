package modules

import (
	ps "backend/internal/modules/pet/service"
	us "backend/internal/modules/user/service"
	"backend/internal/repository"
)

type Services struct {
	Pet  ps.PetServicer
	User us.UserServicer
}

func NewServices(db repository.Repository) *Services {
	return &Services{
		Pet:  ps.NewPetService(db),
		User: us.NewUserService(db),
	}
}
