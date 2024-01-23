package helpers

import (
	"projects/LDmitryLD/petstore/app/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UpdateUser(t *testing.T) {
	slc := []*models.User{{ID: 1, FirstName: "OldName"}, {ID: 2, FirstName: "JustName"}}

	expect := []*models.User{{ID: 1, FirstName: "OldName"}, {ID: 2, FirstName: "WantName"}}

	UpdateUser(&slc, *expect[1])

	assert.Equal(t, expect, slc)
}

func Test_DeleteUser(t *testing.T) {
	slc := []*models.User{{ID: 1, Username: "OldName"}, {ID: 2, Username: "JustName"}}

	expect := []*models.User{{ID: 1, Username: "OldName"}}

	DeleteUser(&slc, "JustName")

	assert.Equal(t, expect, slc)
}
