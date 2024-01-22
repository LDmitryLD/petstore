package modules

import (
	petservice "projects/LDmitryLD/petstore/app/internal/modules/pet/service"
	userservice "projects/LDmitryLD/petstore/app/internal/modules/user/service"
	"projects/LDmitryLD/petstore/app/internal/storages"
)

type Services struct {
	Pet  petservice.Peterer
	User userservice.Userer
}

func NewServices(storages *storages.Storeges) *Services {
	petService := petservice.NewPetService(storages.Pet)
	userService := userservice.NewUserService(storages.User)

	return &Services{
		Pet:  petService,
		User: userService,
	}
}
