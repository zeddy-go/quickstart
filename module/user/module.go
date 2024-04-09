package user

import (
	"github.com/zeddy-go/zeddy/app"
	"github.com/zeddy-go/zeddy/container"
	"github.com/zeddy-go/zeddy/httpx/ginx"
	"quickstart/module/user/domain"
	"quickstart/module/user/iface/http"
	"quickstart/module/user/infra/migration"
	"quickstart/module/user/infra/repo"
)

func NewModule() *Module {
	m := &Module{}

	return m
}

type Module struct {
	app.IsModule
}

func (m Module) Init() (err error) {
	err = container.Bind[*http.UserHandler](http.NewUserHandler)
	if err != nil {
		return
	}

	err = container.Invoke(migration.RegisterMigration)
	if err != nil {
		return
	}

	err = container.Bind[domain.IUserRepo](repo.NewUserRepo)
	if err != nil {
		return
	}

	return
}

func (m Module) Boot() (err error) {
	err = container.Invoke(func(r ginx.Router, userHandler *http.UserHandler) {
		r.GET("/hello", userHandler.Hello)
		r.POST("/api/users", userHandler.Create)
		r.GET("/api/users/:id", userHandler.Detail)
	})
	if err != nil {
		return
	}

	return
}
