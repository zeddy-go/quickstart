package migration

import (
	"embed"
)

//go:embed *.sql
var MigrateFs embed.FS
