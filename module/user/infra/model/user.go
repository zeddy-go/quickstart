package model

import "github.com/zeddy-go/zeddy/database/gormx"

type User struct {
	gormx.CommonField
	Username string
	Password string
}
