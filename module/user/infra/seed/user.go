package seed

import (
	"errors"
	"gorm.io/gorm"
	"quickstart/module/user/domain"
)

func SeedUser(userRepo domain.UserRepo) (err error) {
	_, err = userRepo.First([]any{"username", "test"})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}

		return userRepo.Create(&domain.User{
			Username: "test",
			Password: "test",
		})
	}

	return
}
