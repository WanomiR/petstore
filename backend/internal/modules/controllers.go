package modules

import (
	"backend/internal/lib/rr"
	pc "backend/internal/modules/pet/controller"
	uc "backend/internal/modules/user/controller"
)

type Controllers struct {
	Pet  pc.PetController
	User uc.UserController
}

func NewControllers(services *Services, readResponder rr.ReadResponder) *Controllers {
	return &Controllers{
		Pet:  pc.NewPetControl(services.Pet, readResponder),
		User: uc.NewUserController(services.User, readResponder),
	}
}
