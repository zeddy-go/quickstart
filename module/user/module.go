package user

import (
	"github.com/zeddy-go/zeddy/container"
	"github.com/zeddy-go/zeddy/contract"
	"github.com/zeddy-go/zeddy/module"
	"quickstart/module/user/iface/http"
)

func NewModule() *Module {
	m := &Module{}

	return m
}

type Module struct {
	module.BaseModule
}

func (m Module) Init() (err error) {
	err = container.Bind[*http.UserHandler](http.NewUserHandler)
	if err != nil {
		return
	}

	return
}

func (m Module) Boot() (err error) {
	err = container.Invoke(func(r contract.IRouter, userHandler *http.UserHandler) {
		r.GET("/hello", userHandler.Hello)
	})
	if err != nil {
		return
	}

	return
}
