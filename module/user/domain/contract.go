package domain

import (
	"github.com/zeddy-go/zeddy/database"
	"gorm.io/gorm"
)

type IUserRepo interface {
	database.IRepository[User, *gorm.DB]
}
