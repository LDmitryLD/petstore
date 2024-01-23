package storages

import (
	petstorage "projects/LDmitryLD/petstore/app/internal/modules/pet/storage"
	storestorage "projects/LDmitryLD/petstore/app/internal/modules/store/storage"
	userstorage "projects/LDmitryLD/petstore/app/internal/modules/user/storage"
)

type Storeges struct {
	Pet   petstorage.PetStorager
	Store storestorage.StoreStorager
	User  userstorage.UserStorager
}

func NewStorages() *Storeges {
	return &Storeges{
		Pet:   petstorage.NewPetStorage(),
		Store: storestorage.NewStorage(),
		User:  userstorage.NewUserStorage(),
	}
}
