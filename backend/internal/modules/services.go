package modules

import (
	ps "backend/internal/modules/pet/service"
	"backend/internal/repository"
)

type Services struct {
	Pet ps.PetServicer
}

func NewServices(db repository.Repository) *Services {
	return &Services{
		Pet: ps.NewPetService(db),
	}
}
