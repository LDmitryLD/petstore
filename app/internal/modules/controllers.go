package modules

import (
	authcontroller "projects/LDmitryLD/petstore/app/internal/modules/auth/controller"
	petcontroller "projects/LDmitryLD/petstore/app/internal/modules/pet/controller"
	storecontroller "projects/LDmitryLD/petstore/app/internal/modules/store/controller"
	usercontroller "projects/LDmitryLD/petstore/app/internal/modules/user/controller"
)

type Controllers struct {
	Auth  authcontroller.Auther
	Pet   petcontroller.Peterer
	Store storecontroller.Storeger
	User  usercontroller.Userer
}

func NewControllers(services *Services) *Controllers {
	authController := authcontroller.NewAuth(services.Auth)
	petController := petcontroller.NewPet(services.Pet)
	userController := usercontroller.NewUser(services.User)
	storeController := storecontroller.NewStore(services.Store)

	return &Controllers{
		Auth:  authController,
		Pet:   petController,
		Store: storeController,
		User:  userController,
	}
}
