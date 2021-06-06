package bot

import (
	"fmt"
	"reflect"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/sophistik/ITDigestBot/internal/appconfig"
	"github.com/sophistik/ITDigestBot/internal/services"
)

const (
	addCommand    string = "add"
	getCommand    string = "get"
	deleteCommand string = "delete"
)

type Bot struct {
	cfg appconfig.Bot
	bot *tgbotapi.BotAPI

	botApiService *services.BotAPIService
}

func NewBot(
	cfg appconfig.Bot,
	bas *services.BotAPIService,
) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		bot:           bot,
		botApiService: bas,
	}, err
}

func (b *Bot) Run() error {
	var (
		msg  tgbotapi.MessageConfig
		ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	)

	b.bot.Debug = true
	ucfg.Timeout = b.cfg.Timeout
	updates, err := b.bot.GetUpdatesChan(ucfg)
	if err != nil {
		return fmt.Errorf("can't create updates chan: %w", err)
	}

	for u := range updates {
		if u.Message == nil { // ignore any non-Message Updates
			continue
		}

		if validMessage := reflect.TypeOf(u.Message.Text).Kind() == reflect.String && u.Message.Text != ""; !validMessage {
			msg = tgbotapi.NewMessage(u.Message.Chat.ID, "Невалидноее сообщение")
			b.bot.Send(msg)

			continue
		}

		if u.Message.Command() != "" {
			b.processCommand(u.Message)
			continue
		}

		b.processMessage(u.Message)
	}

	return nil
}

func (b *Bot) processCommand(m *tgbotapi.Message) {
	var msg tgbotapi.MessageConfig

	switch m.Command() {
	case addCommand:
		msg = tgbotapi.NewMessage(m.Chat.ID, "Введите теги через запятую, например:\nC++, Golang")
		b.botApiService.SetLastInput(m.Chat.ID, addCommand)
	case deleteCommand:
		msg = tgbotapi.NewMessage(m.Chat.ID, "Введите теги через запятую, например:\nC++, Golang")
		b.botApiService.SetLastInput(m.Chat.ID, deleteCommand)
	case getCommand:
		tags, _ := b.botApiService.GetTags(m.Chat.ID)

		msg = tgbotapi.NewMessage(m.Chat.ID, strings.Join(tags, ", "))
	}

	b.bot.Send(msg)
}

func (b *Bot) processMessage(m *tgbotapi.Message) {
	var msg tgbotapi.MessageConfig

	lastInput, err := b.botApiService.GetLastInput(m.Chat.ID)
	if err != nil {
		msg = tgbotapi.NewMessage(m.Chat.ID, "Сначала введите команду")
		b.bot.Send(msg)

		return
	}

	switch lastInput {
	case addCommand:
		tags := strings.Split(m.Text, ", ")
		err = b.botApiService.AddTags(m.Chat.ID, tags)

	case deleteCommand:
		tags := strings.Split(m.Text, ", ")
		err = b.botApiService.RemoveTags(m.Chat.ID, tags)
	}

	msg = tgbotapi.NewMessage(m.Chat.ID, "Готово!")
	if err != nil {
		msg = tgbotapi.NewMessage(m.Chat.ID, "Кажется, что-то пошло не так, попробуйте повторить попытку позже.")
	}

	b.botApiService.RemoveLastUpdate(m.Chat.ID)

	b.bot.Send(msg)
}
