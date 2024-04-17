package main

import (
	"github.com/zeddy-go/zeddy/app"
	"github.com/zeddy-go/zeddy/configx"
	"github.com/zeddy-go/zeddy/database/migrate"
	"github.com/zeddy-go/zeddy/database/wgorm"
	"github.com/zeddy-go/zeddy/httpx/ginx"
	"github.com/zeddy-go/zeddy/httpx/grpcx"
	"log/slog"
	"quickstart/conf"
	"quickstart/module/user"
)

func main() {
	app.Use(
		configx.NewModule(configx.WithContent(conf.Config)),
		wgorm.NewModule(),
		migrate.NewModule(),
		ginx.NewModule(),
		grpcx.NewModule(),
		user.NewModule(),
	)
	err := app.StartAndWait()
	if err != nil {
		slog.Warn("An error occurred", "error", err)
	}
}
