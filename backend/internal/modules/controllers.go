package modules

import (
	"backend/internal/lib/rr"
	pc "backend/internal/modules/pet/controller"
	sc "backend/internal/modules/store/controller"
	uc "backend/internal/modules/user/controller"
)

type Controllers struct {
	Pet   pc.PetController
	User  uc.UserController
	Store sc.StoreController
}

func NewControllers(services *Services, readResponder rr.ReadResponder) *Controllers {
	return &Controllers{
		Pet:   pc.NewPetControl(services.Pet, readResponder),
		User:  uc.NewUserController(services.User, readResponder),
		Store: sc.NewStoreControl(services.Store, readResponder),
	}
}
