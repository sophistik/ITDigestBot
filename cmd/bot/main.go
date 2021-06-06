package main

import (
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/sophistik/ITDigestBot/internal/appconfig"
	"github.com/sophistik/ITDigestBot/internal/bot"
	"github.com/sophistik/ITDigestBot/internal/postgres"
	"github.com/sophistik/ITDigestBot/internal/services"
	pg "github.com/sophistik/ITDigestBot/pkg/postgres"
	"go.uber.org/zap"
)

func parseConfig() appconfig.BotAPI {
	var cfg appconfig.BotAPI

	cfg.Bot.Token = "ТОКЕН"
	cfg.Bot.Timeout = 60

	cfg.Postgres.User = "digestbot"
	cfg.Postgres.Password = "digestbot"
	cfg.Postgres.Database = "digestbot"
	cfg.Postgres.Host = "localhost"
	cfg.Postgres.Port = 5432
	cfg.Postgres.Params = "?sslmode=disable&fallback_application_name=bot_api"
	cfg.Postgres.MaxConns = 5
	cfg.Postgres.MaxConnLifetime = time.Hour

	return cfg
}

func main() {
	cfg := parseConfig()

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("logger")
	}

	db, err := pg.NewDB(logger, cfg.Postgres)
	if err != nil {
		logger.Sugar().Fatal("can't create db: ", err)
	}

	goquDB := goqu.New("postgres", db.Session)

	userStorage, err := postgres.NewUserTagsStorage(db, goquDB)
	if err != nil {
		logger.Sugar().Fatal("can't create user storage: ", err)
	}

	botApiService := services.NewBotAPIService(userStorage)

	bot, err := bot.NewBot(cfg.Bot, botApiService)
	if err != nil {
		logger.Sugar().Fatal("can't create bot API: ", err)
	}

	err = bot.Run()
	if err != nil {
		logger.Sugar().Fatal("can't run API: ", err)
	}

}
