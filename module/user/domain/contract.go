package domain

import (
	"github.com/zeddy-go/zeddy/database"
	"gorm.io/gorm"
)

type UserRepo interface {
	database.IRepository[User, *gorm.DB]
}
