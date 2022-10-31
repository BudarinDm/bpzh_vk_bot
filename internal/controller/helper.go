package controller

import (
	"bpzh_vk_bot/internal/model/domain"
	"context"
	"encoding/json"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
)

func (a *App) botHandler(obj events.MessageNewObject) error {
	b := params.NewMessagesSendBuilder()
	keyboards := domain.Keyboard{
		OneTime: true,
		Buttons: [][]domain.Button{
			{
				{
					Action: domain.Action{
						Type:    "callback",
						Payload: `{"command": "all-menu"}`,
						Label:   "Тегнуть группу",
					},
					Color: "primary",
				},
			},
		},
	}

	by, _ := json.Marshal(keyboards)
	b.Keyboard(string(by))

	b.Message("Чем могу помочь?")
	b.RandomID(0)
	b.PeerID(obj.Message.PeerID)

	st, err := a.vk.MessagesSend(b.Params)
	fmt.Println(st)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) allMenuHandler(obj events.MessageEventObject) error {
	b := params.NewMessagesSendBuilder()

	var buttons [][]domain.Button
	count := 0

	groups, err := a.logic.GetGroupsByChatId(context.Background(), int64(obj.PeerID))
	if err != nil {
		return err
	}

	if len(groups) != 0 {
		var arr []domain.Button

		for _, g := range groups {
			arr = append(arr,
				domain.Button{
					Action: domain.Action{
						Type:    "callback",
						Payload: fmt.Sprintf(`{"command": "all", "arg": "%s"}`, g.Name),
						Label:   g.Name,
					},
					Color: g.Color,
				})
			count += 1
			if count == 3 {
				buttons = append(buttons, arr)
				arr = []domain.Button{}
				count = 0
			}
		}
		if len(arr) > 0 {
			buttons = append(buttons, arr)
		}

		keyboards := domain.Keyboard{
			OneTime: true,
			Buttons: buttons,
		}

		by, _ := json.Marshal(keyboards)
		b.Keyboard(string(by))

		b.Message("Выберите группу")
		b.RandomID(0)
		b.PeerID(obj.PeerID)

		_, err = a.vk.MessagesSend(b.Params)
		if err != nil {
			return err
		}

		return nil
	}

	b.Message("Группы не созданы, для добавления групп пользователей воспользуйтесь /group info")
	b.RandomID(0)
	b.PeerID(obj.PeerID)

	st, err := a.vk.MessagesSend(b.Params)
	fmt.Println(st)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) nextStepHandler(obj events.MessageEventObject) error {
	b := params.NewMessagesSendBuilder()
	var msg string
	msg = "Вторая часть меню"

	keyboards := domain.Keyboard{
		OneTime: true,
		Buttons: [][]domain.Button{
			{
				{
					Action: domain.Action{
						Type:    "text",
						Payload: `{"command": "all", "arg": "voronezh"}`,
						Label:   "In developing...",
					},
					Color: "primary",
				},
				{
					Action: domain.Action{
						Type:    "text",
						Payload: `{"command": "all", "arg": "moscow"}`,
						Label:   "In developing...",
					},
					Color: "primary",
				},
			},
		},
	}

	by, _ := json.Marshal(keyboards)
	b.Keyboard(string(by))

	b.Message(msg)
	b.RandomID(0)
	b.PeerID(obj.PeerID)
	_, err := a.vk.MessagesSend(b.Params)
	if err != nil {
		return err
	}

	return nil
}
