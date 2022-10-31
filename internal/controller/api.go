package controller

import (
	"bpzh_vk_bot/internal/config"
	"bpzh_vk_bot/internal/logic"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"github.com/rs/zerolog/log"
)

// App основная структура для приложения
type App struct {
	lp     *longpoll.LongPoll
	vk     *api.VK
	config *config.Config
	logic  *logic.Logic
}

func NewApp(config *config.Config, logic *logic.Logic, longPoll *longpoll.LongPoll, vk *api.VK) *App {
	return &App{
		vk:     vk,
		lp:     longPoll,
		config: config,
		logic:  logic,
	}
}

func (a *App) StartServe() {

	a.handler()
	// Run Bots Long Poll
	log.Info().Msg("Start Long Poll")
	if err := a.lp.Run(); err != nil {
		log.Error().Err(err)
	}
}

type MessageEventResponse struct {
	Command string `json:"command"`
	Arg     string `json:"arg"`
}
