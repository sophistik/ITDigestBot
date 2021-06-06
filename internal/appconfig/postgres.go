package appconfig

import "time"

type DB struct {
	URL             string
	MaxConnections  int
	MaxConnLifetime time.Duration
}
