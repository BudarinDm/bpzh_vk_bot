package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/rs/zerolog/log"
	"strings"
)

func (a *App) handler() {
	a.lp.MessageEvent(func(_ context.Context, obj events.MessageEventObject) {
		log.Printf("%d: %s, %s, %d, %d", obj.UserID, obj.EventID, obj.Payload, obj.PeerID, obj.ConversationMessageID)

		var e MessageEventResponse
		err := json.Unmarshal(obj.Payload, &e)
		if err != nil {
			log.Error().Err(err).Msg("handler Unmarshal")
			err = a.sendMsgEventBuilder(&obj, err.Error())
			if err != nil {
				return
			}
		}

		if e.Command == "all-menu" {
			err = a.allMenuHandler(obj)
			if err != nil {
				log.Error().Err(err).Msg("all-menu")
				err = a.sendMsgEventBuilder(&obj, err.Error())
				if err != nil {
					return
				}
			}
		}
	})

	a.lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
		log.Printf("%d: %s. authorId: %d", obj.Message.PeerID, obj.Message.Text, obj.Message.FromID)

		msg := obj.Message.Text
		splitMsgs := strings.Split(msg, " ")

		if msg == "/settings" && obj.Message.PeerID == 144568579 {
			fmt.Println("/settings")
		}

		if splitMsgs[0] == "/group" {
			err := a.groupRouter(splitMsgs, obj)
			if err != nil {
				log.Error().Err(err).Msg("/group")
				err = a.sendMsgBuilder(&obj, err.Error())
				if err != nil {
					return
				}
			}
		}

		if msg == "/bot" {
			err := a.botHandler(obj)
			if err != nil {
				log.Error().Err(err)
			}
		}
	})
}

func (a *App) sendMsgBuilder(obj *events.MessageNewObject, msg string) error {
	b := params.NewMessagesSendBuilder()

	b.Message(msg)
	b.RandomID(0)
	b.PeerID(obj.Message.PeerID)

	st, err := a.vk.MessagesSend(b.Params)
	fmt.Println(st)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) sendMsgEventBuilder(obj *events.MessageEventObject, msg string) error {
	b := params.NewMessagesSendBuilder()

	b.Message(msg)
	b.RandomID(0)
	b.PeerID(obj.PeerID)

	st, err := a.vk.MessagesSend(b.Params)
	fmt.Println(st)
	if err != nil {
		return err
	}
	return nil
}
