package storage

import (
	"fmt"
	"projects/LDmitryLD/petstore/app/helpers"
	"projects/LDmitryLD/petstore/app/internal/models"
	"sync"
)

type UserStorage struct {
	users              []*models.User
	userNameKeys       map[string]*models.User
	autoIncrementCount int
	sync.Mutex
}

func NewUserStorage() *UserStorage {
	return &UserStorage{
		users:              make([]*models.User, 0, 13),
		userNameKeys:       make(map[string]*models.User, 13),
		autoIncrementCount: 1,
	}
}

func (u *UserStorage) Create(user models.User) int {
	u.Lock()
	defer u.Unlock()

	user.ID = u.autoIncrementCount
	u.userNameKeys[user.Username] = &user
	u.autoIncrementCount++
	u.users = append(u.users, &user)

	return user.ID
}

func (u *UserStorage) GetByUsername(userName string) (models.User, error) {
	u.Lock()
	defer u.Unlock()

	if u, ok := u.userNameKeys[userName]; ok {
		return *u, nil
	}

	return models.User{}, fmt.Errorf("not found")
}

func (u *UserStorage) Update(userName string, user models.User) error {
	u.Lock()
	defer u.Unlock()

	if _, ok := u.userNameKeys[userName]; !ok {
		return fmt.Errorf("not found")
	}

	u.userNameKeys[userName] = &user

	helpers.UpdateUser(&u.users, user)

	return nil
}

func (u *UserStorage) Delete(userName string) error {
	u.Lock()
	defer u.Unlock()

	if _, ok := u.userNameKeys[userName]; !ok {
		return fmt.Errorf("not found")
	}

	helpers.DeleteUser(&u.users, userName)

	delete(u.userNameKeys, userName)

	return nil
}
