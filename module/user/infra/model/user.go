package model

import "github.com/zeddy-go/zeddy/database/wgorm"

type User struct {
	wgorm.CommonField
	Username string
	Password string
}
