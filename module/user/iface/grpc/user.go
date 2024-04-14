package grpc

import (
	"context"
	"github.com/zeddy-go/zeddy/errx"
	"github.com/zeddy-go/zeddy/mapper"
	"quickstart/module/user/domain"
	"quickstart/module/user/domain/svc"
	"quickstart/pb"
)

func NewUserSvc(userSvc *svc.User) *UserSrv {
	return &UserSrv{userSvc: userSvc}
}

type UserSrv struct {
	pb.UnimplementedUserSvcServer
	userSvc *svc.User
}

func (u UserSrv) Create(ctx context.Context, req *pb.User) (resp *pb.ID, err error) {
	var user = &domain.User{}
	err = mapper.SimpleMap(user, req)
	if err != nil {
		return
	}

	err = u.userSvc.Create(user)
	if err != nil {
		err = errx.Wrap(err, "create user failed")
		return
	}

	resp = &pb.ID{
		Id: user.ID,
	}
	return
}

func (u UserSrv) Detail(ctx context.Context, req *pb.ID) (resp *pb.User, err error) {
	user, err := u.userSvc.Detail(req.Id)
	if err != nil {
		err = errx.Wrap(err, "find user failed")
		return
	}

	resp = &pb.User{}
	err = mapper.SimpleMap(resp, user)
	return
}
