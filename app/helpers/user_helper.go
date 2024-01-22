package helpers

import "projects/LDmitryLD/petstore/app/internal/models"

func UpdateUser(users *[]*models.User, user models.User) {
	for i, u := range *users {
		if u.ID == user.ID {
			(*users)[i] = &user
		}
	}
}

func DeleteUser(users *[]*models.User, userName string) {
	for i, u := range *users {
		if u.Username == userName {
			*users = append((*users)[:i], (*users)[i+1:]...)
			break
		}
	}
}
