package modules

import (
	petcontroller "projects/LDmitryLD/petstore/app/internal/modules/pet/controller"
	usercontroller "projects/LDmitryLD/petstore/app/internal/modules/user/controller"
)

type Controllers struct {
	Pet  petcontroller.Peterer
	User usercontroller.Userer
}

func NewControllers(services *Services) *Controllers {
	petController := petcontroller.NewPet(services.Pet)
	userController := usercontroller.NewUser(services.User)

	return &Controllers{
		Pet:  petController,
		User: userController,
	}
}
