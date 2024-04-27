package repo

import (
	"github.com/zeddy-go/zeddy/database"
	"github.com/zeddy-go/zeddy/database/gormx"
	"gorm.io/gorm"
	"quickstart/module/user/domain"
	"quickstart/module/user/infra/model"
)

func NewUserRepo(db *gorm.DB) domain.UserRepo {
	r := &UserRepository{}
	r.IRepository = gormx.NewRepository[model.User, domain.User](db)
	return r
}

type UserRepository struct {
	database.IRepository[domain.User]
}
