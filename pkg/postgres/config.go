package postgres

import (
	"fmt"
	"time"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     uint
	Database string
	Params   string

	MaxConns        int
	MaxConnLifetime time.Duration
}

func (c Config) DNS() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s%s", c.User, c.Password, c.Host, c.Port, c.Database, c.Params)
}
