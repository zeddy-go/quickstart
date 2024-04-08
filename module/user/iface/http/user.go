package http

import (
	"github.com/zeddy-go/zeddy/database"
	"quickstart/module/user/domain"
)

func NewUserHandler(userRepo domain.IUserRepo) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

type UserHandler struct {
	userRepo domain.IUserRepo
}

func (uh *UserHandler) Hello(req *HelloReq) *HelloResp {
	return &HelloResp{
		Text: "hello " + req.Username + "!",
	}
}

type HelloReq struct {
	Username string `form:"username" binding:"required"`
}

type HelloResp struct {
	Text string `json:"text"`
}

func (uh *UserHandler) Create(req *CreateReq) (resp *CreateResp, err error) {
	user := &domain.User{
		Username: req.Username,
		Password: req.Password,
	}
	err = uh.userRepo.Create(user)
	if err != nil {
		return
	}
	resp = &CreateResp{
		ID: user.ID,
	}
	return
}

type CreateReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateResp struct {
	ID uint64 `json:"id,string"`
}

func (uh *UserHandler) Detail(req *DetailReq) (resp *DetailResp, err error) {
	return uh.userRepo.First(database.Condition{"id", req.ID})
}

type DetailReq struct {
	ID uint64 `uri:"id" binding:"required"`
}

type DetailResp = domain.User
