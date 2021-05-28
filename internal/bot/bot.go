package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type Bot struct {
	Bot *tgbotapi.BotAPI
}

func NewBot(token string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		Bot: bot,
	}, err
}
