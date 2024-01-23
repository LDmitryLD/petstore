package modules

import (
	authservice "projects/LDmitryLD/petstore/app/internal/modules/auth/service"
	petservice "projects/LDmitryLD/petstore/app/internal/modules/pet/service"
	storeservice "projects/LDmitryLD/petstore/app/internal/modules/store/service"
	userservice "projects/LDmitryLD/petstore/app/internal/modules/user/service"
	"projects/LDmitryLD/petstore/app/internal/storages"
)

type Services struct {
	Auth  authservice.Auther
	Pet   petservice.Peterer
	Store storeservice.Storeger
	User  userservice.Userer
}

func NewServices(storages *storages.Storeges) *Services {
	petService := petservice.NewPetService(storages.Pet)
	userService := userservice.NewUserService(storages.User)
	storeService := storeservice.NewStoreService(storages.Store)
	authservice := authservice.NewAuth()

	return &Services{
		Auth:  authservice,
		Pet:   petService,
		Store: storeService,
		User:  userService,
	}
}
