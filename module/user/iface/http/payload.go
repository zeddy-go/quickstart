package http

import "quickstart/module/user/domain"

type HelloReq struct {
	Username string `form:"username" binding:"required"`
}

type HelloResp struct {
	Text string `json:"text"`
}
type CreateReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateResp struct {
	ID uint64 `json:"id,string"`
}
type DetailReq struct {
	ID uint64 `uri:"id" binding:"required"`
}

type DetailResp = domain.User
