package migration

import (
	"embed"
	"github.com/zeddy-go/zeddy/database/migrate"
)

//go:embed *.sql
var migrateFs embed.FS

func RegisterMigration(driver *migrate.EmbedDriver) {
	driver.Add(migrateFs)
}
