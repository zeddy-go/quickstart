package http

import (
	"quickstart/module/user/domain"
	"quickstart/module/user/domain/svc"
)

func NewUserHandler(userSvc *svc.User) *UserHandler {
	return &UserHandler{
		userSvc: userSvc,
	}
}

type UserHandler struct {
	userSvc *svc.User
}

func (uh *UserHandler) Hello(req *HelloReq) *HelloResp {
	return &HelloResp{
		Text: "hello " + req.Username + "!",
	}
}

func (uh *UserHandler) Create(req *CreateReq) (resp *CreateResp, err error) {
	user := &domain.User{
		Username: req.Username,
		Password: req.Password,
	}
	err = uh.userSvc.Create(user)
	if err != nil {
		return
	}
	resp = &CreateResp{
		ID: user.ID,
	}
	return
}

func (uh *UserHandler) Detail(req *DetailReq) (resp *DetailResp, err error) {
	return uh.userSvc.Detail(req.ID)
}
