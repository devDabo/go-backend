package sqlmigrations

import "embed"

// FS contains SQL migration files embedded at build time.
//
//go:embed migrations/*.sql
var FS embed.FS
