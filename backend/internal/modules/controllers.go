package modules

import (
	"backend/internal/lib/rr"
	pc "backend/internal/modules/pet/controller"
)

type Controllers struct {
	Pet pc.PetController
}

func NewControllers(services *Services, readResponder rr.ReadResponder) *Controllers {
	return &Controllers{
		Pet: pc.NewPetControl(services.Pet, readResponder),
	}
}
