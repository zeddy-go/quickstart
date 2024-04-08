package http

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

type UserHandler struct{}

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
