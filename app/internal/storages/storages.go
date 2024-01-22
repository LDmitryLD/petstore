package storages

import (
	petstorage "projects/LDmitryLD/petstore/app/internal/modules/pet/storage"
	userstorage "projects/LDmitryLD/petstore/app/internal/modules/user/storage"
)

type Storeges struct {
	Pet  petstorage.PetStorager
	User userstorage.UserStorager
}

func NewStorages() *Storeges {
	return &Storeges{
		Pet:  petstorage.NewPetStorage(),
		User: userstorage.NewUserStorage(),
	}
}
