package poller

import (
	tapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Poller struct {
	bots []*tapi.BotAPI
}

func New() (p *Poller, err error) {
	return
}

func (p *Poller) Run() {

}
