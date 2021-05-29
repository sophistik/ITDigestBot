package main

import (
	"fmt"

	"github.com/sophistik/ITDigestBot/internal/appconfig"
	"github.com/sophistik/ITDigestBot/internal/bot"
	"github.com/sophistik/ITDigestBot/internal/storage"
)

func parseConfig() appconfig.BotAPI {
	var cfg appconfig.BotAPI

	cfg.Bot.Token = "Token"
	cfg.Bot.Timeout = 60

	return cfg
}

func main() {
	cfg := parseConfig()

	userTagsStorage, err := storage.NewUserTagsStorage()
	if err != nil {
		fmt.Println("can't create user tags storage: ", err)
		return
	}

	userInputStorage, err := storage.NewLastUserInputStorage()
	if err != nil {
		fmt.Println("can't create user input storage: ", err)
		return
	}

	bot, err := bot.NewBot(cfg.Bot, userTagsStorage, userInputStorage)
	if err != nil {
		// fmt.Errorf("err: %w", err)
		fmt.Println("can't create bot: ", err)
		return
	}

	err = bot.Run()
	if err != nil {
		// fmt.Errorf("cant't run bot: %w", err)
		fmt.Println("can't run bot", err)
		return
	}

}
