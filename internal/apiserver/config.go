package apiserver

import "github.com/aabbuukkaarr8/TG-BOT/internal/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}

//psql "postgres://appuser:secret@localhost:5432/erziapp?sslmode=disable"

//psql "postgres://appuser:secret@localhost:5432/erziapp?sslmode=disable" \
//-c "ALTER TABLE products ADD COLUMN category TEXT NOT NULL DEFAULT 'general';"
