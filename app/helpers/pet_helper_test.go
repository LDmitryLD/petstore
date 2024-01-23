package helpers

import (
	"projects/LDmitryLD/petstore/app/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DeletePet(t *testing.T) {
	slc := []*models.Pet{{ID: 1}, {ID: 2}}

	DeletePet(&slc, 2)

	expect := []*models.Pet{{ID: 1}}

	assert.Equal(t, expect, slc)
}

func Test_UpdatePet(t *testing.T) {
	slc := []*models.Pet{{ID: 1, Name: "DOG"}}

	expect := []*models.Pet{{ID: 1, Name: "CAT"}}

	UpdatePet(&slc, *expect[0])

	assert.Equal(t, expect, slc)
}

func Test_UpdatePetByID(t *testing.T) {
	slc := []*models.Pet{{ID: 1, Name: "DOG", Status: "A"}, {ID: 2, Name: "CAT", Status: "B"}}

	expect := []*models.Pet{{ID: 1, Name: "DOG", Status: "A"}, {ID: 2, Name: "DOG2", Status: "C"}}

	UpdatePetByID(&slc, 2, "DOG2", "C")

	assert.Equal(t, expect, slc)
}
