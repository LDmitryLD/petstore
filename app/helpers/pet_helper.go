package helpers

import (
	"projects/LDmitryLD/petstore/app/internal/models"
)

func DeletePet(pets *[]*models.Pet, id int) {
	for i, pet := range *pets {
		if pet.ID == id {
			*pets = append((*pets)[:i], (*pets)[i+1:]...)
			break
		}
	}
}

func UpdatePet(pets *[]*models.Pet, pet models.Pet) {
	for i, p := range *pets {
		if p.ID == pet.ID {
			(*pets)[i] = &pet
		}
	}
}

func UpdatePetByID(pets *[]*models.Pet, id int, name string, status string) {
	for i, p := range *pets {
		if p.ID == id {
			switch {
			case name != "":
				(*pets)[i].Name = name
				fallthrough
			case status != "":
				(*pets)[i].Status = status
			}
		}
	}
}
