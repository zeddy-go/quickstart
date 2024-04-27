package domain

import (
	"github.com/zeddy-go/zeddy/database"
)

type UserRepo interface {
	database.IRepository[User]
}
