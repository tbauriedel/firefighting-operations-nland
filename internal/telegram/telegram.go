package telegram

import (
	"fmt"
	t "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

type Bot struct {
	Instance *t.BotAPI
}

func NewBotInstance(token string) (Bot, error) {
	b := Bot{}
	var err error

	b.Instance, err = t.NewBotAPI(token)
	if err != nil {
		return Bot{}, fmt.Errorf("cant create bot api instance: %w", err)
	}

	return b, nil
}

func (b *Bot) Send(chatId int64, message string) error {
	m := t.NewMessage(chatId, message)
	if _, err := b.Instance.Send(m); err != nil {
		return err
	}

	log.Printf("new operation has been sent. Message: %s", strings.Replace(message, "\n", "\\n", -1))

	return nil
}
