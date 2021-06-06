package appconfig

import "github.com/sophistik/ITDigestBot/pkg/postgres"

type BotAPI struct {
	Bot
	Postgres postgres.Config
}

type Bot struct {
	Token   string
	Timeout int
}
