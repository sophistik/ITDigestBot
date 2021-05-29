package bot

import (
	"fmt"
	"reflect"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/sophistik/ITDigestBot/internal/appconfig"
	"github.com/sophistik/ITDigestBot/internal/repos"
)

const (
	addCommand string = "add"
	getCommand string = "get"
)

type Bot struct {
	Cfg appconfig.Bot
	Bot *tgbotapi.BotAPI

	UserTagsStorage   repos.UserTagsRepo
	LastInputsStorage repos.LastUserInput
}

func NewBot(
	cfg appconfig.Bot,
	uts repos.UserTagsRepo,
	lis repos.LastUserInput,
) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		Bot:               bot,
		UserTagsStorage:   uts,
		LastInputsStorage: lis,
	}, err
}

func (b *Bot) Run() error {
	var (
		msg  tgbotapi.MessageConfig
		ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	)

	b.Bot.Debug = true
	ucfg.Timeout = b.Cfg.Timeout
	updates, err := b.Bot.GetUpdatesChan(ucfg)
	if err != nil {
		return fmt.Errorf("can't create updates chan: %w", err)
	}

	for u := range updates {
		if u.Message == nil { // ignore any non-Message Updates
			continue
		}

		if validMessage := reflect.TypeOf(u.Message.Text).Kind() == reflect.String && u.Message.Text != ""; !validMessage {
			msg = tgbotapi.NewMessage(u.Message.Chat.ID, "Невалидноее сообщение.")
			b.Bot.Send(msg)

			continue
		}

		if u.Message.Command() != "" {
			b.processCommand(u.Message)
			continue
		}

		lastInput, err := b.LastInputsStorage.Get(u.Message.Chat.ID)
		if err != nil {
			msg = tgbotapi.NewMessage(u.Message.Chat.ID, "Сначала введите команду.")
			b.Bot.Send(msg)

			continue
		}

		switch lastInput {
		case addCommand:
			tags := strings.Split(u.Message.Text, ", ")
			b.UserTagsStorage.Upsert(u.Message.Chat.ID, tags)

			b.LastInputsStorage.Delete(u.Message.Chat.ID)
		}

		b.Bot.Send(msg)
	}

	return nil
}

func (b *Bot) processCommand(m *tgbotapi.Message) {
	var msg tgbotapi.MessageConfig

	switch m.Command() {
	case addCommand:
		msg = tgbotapi.NewMessage(m.Chat.ID, "Введите теги через запятую, например:\nC++, Golang")
		b.LastInputsStorage.Add(m.Chat.ID, addCommand)
	case getCommand:
		tags, _ := b.UserTagsStorage.Get(m.Chat.ID)

		msg = tgbotapi.NewMessage(m.Chat.ID, strings.Join(tags, ", "))
	}

	b.Bot.Send(msg)
}
