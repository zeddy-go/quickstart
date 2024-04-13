package svc

import (
	"quickstart/module/user/domain"
)

func NewUserService(repo domain.UserRepo) *User {
	return &User{
		userRepo: repo,
	}
}

type User struct {
	userRepo domain.UserRepo
}

func (uh *User) Create(user *domain.User) (err error) {
	return uh.userRepo.Create(user)
}

func (uh *User) Detail(id uint64) (user *domain.User, err error) {
	return uh.userRepo.First([]any{"id", id})
}
