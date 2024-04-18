package main

import (
	"github.com/zeddy-go/zeddy/app"
	"github.com/zeddy-go/zeddy/configx"
	"github.com/zeddy-go/zeddy/database/gormx"
	"github.com/zeddy-go/zeddy/database/migrate"
	"github.com/zeddy-go/zeddy/httpx/ginx"
	"log/slog"
	"quickstart/conf"
	"quickstart/module/user"
)

func main() {
	app.Use(
		configx.NewModule(configx.WithContent(conf.Config)),
		gormx.NewModule(),
		migrate.NewModule(),
		ginx.NewModule(),
		user.NewModule(),
	)
	err := app.StartAndWait()
	if err != nil {
		slog.Warn("An error occurred", "error", err)
	}
}
