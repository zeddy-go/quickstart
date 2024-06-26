package main

import (
	"github.com/zeddy-go/zeddy/app"
	"github.com/zeddy-go/zeddy/httpx/ginx"
	"github.com/zeddy-go/zeddy/config"
	"github.com/zeddy-go/zeddy/database/migrate"
	"github.com/zeddy-go/zeddy/database/wgorm"
	"log/slog"
	conf "quickstart/config"
	"quickstart/module/user"
)

func main() {
	app.Use(
		config.NewModule(config.WithContent(conf.Config)),
		wgorm.NewModule(),
		migrate.NewModule(),
		ginx.NewModule(),
		user.NewModule(),
	)
	err := app.StartAndWait()
	if err != nil {
		slog.Warn("An error occurred", "error", err)
	}
}
