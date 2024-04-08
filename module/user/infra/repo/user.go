package repo

import (
	"github.com/zeddy-go/zeddy/database"
	"github.com/zeddy-go/zeddy/database/wgorm"
	"gorm.io/gorm"
	"quickstart/module/user/domain"
	"quickstart/module/user/infra/model"
)

func NewUserRepo(db *gorm.DB) domain.IUserRepo {
	r := &UserRepository{}
	r.IRepository = wgorm.NewRepository[model.User, domain.User](db)
	return r
}

type UserRepository struct {
	database.IRepository[domain.User, *gorm.DB]
}
