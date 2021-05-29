package appconfig

type BotAPI struct {
	Bot
}

type Bot struct {
	Token   string
	Timeout int
}
