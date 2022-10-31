package main

import (
	"bpzh_vk_bot/internal/config"
	"bpzh_vk_bot/internal/controller"
	"bpzh_vk_bot/internal/logic"
	"bpzh_vk_bot/internal/repo/db"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed read Config")
	}

	vk := api.NewVK(cfg.App.BotToken)

	fs, err := db.CreateFSConnections(&cfg.DB)
	if err != nil {
		log.Error().Err(err)
	}

	group, err := vk.GroupsGetByID(nil)
	if err != nil {
		log.Error().Err(err)
	}

	// Initializing Long Poll
	lp, err := longpoll.NewLongPoll(vk, group[0].ID)
	if err != nil {
		log.Error().Err(err)
	}

	repo := db.NewRepo(fs)                     // model работает с БД и прочими источниками данных
	lgc := logic.NewLogic(cfg, repo, vk)       // logic знает, что делать с model
	app := controller.NewApp(cfg, lgc, lp, vk) // api использует logic для обработки запросов

	app.StartServe()
}
