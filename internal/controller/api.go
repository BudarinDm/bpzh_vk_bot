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
	//members, err := a.vk.MessagesGetConversationMembers(api.Params{"peer_id": 2000000002})
	//if err != nil {
	//	log.Error().Err(err)
	//}
	//for _, m := range members.Items {
	//	fmt.Println(m.MemberID, m.IsAdmin, m.IsOwner)
	//}
	//,13213283,51899876,32329135
	//users, err := a.vk.UsersGet(api.Params{"user_ids": "13213283,51899876,32329135", "fields": "domain"})
	//if err != nil {
	//	log.Error().Err(err)
	//}
	//for _, u := range users {
	//	fmt.Println(u.ID, u.LastName, u.FirstName)
	//	user := domain.User{
	//		FirstName: u.FirstName,
	//		LastName:  u.LastName,
	//		TgId:      "",
	//		VkId:      int64(u.ID),
	//		VkDomain:  u.Domain,
	//	}
	//	err = a.logic.CreateUser(context.Background(), user)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}

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
	UserID  int    `json:"user_id"`
}
