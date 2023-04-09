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
	a.lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
		log.Printf("%d: %s. authorId: %d", obj.Message.PeerID, obj.Message.Text, obj.Message.FromID)

		msg := obj.Message.Text
		splitMsgs := strings.Split(msg, " ")

		if obj.Message.FromID == 58778743 {
			b := params.NewMessagesSendBuilder()

			b.Message("@anthony_club (Антон)")
			b.RandomID(0)
			b.PeerID(obj.Message.PeerID)
			b.Attachment("photo-216838391_457239034")

			_, err := a.vk.MessagesSend(b.Params)
			if err != nil {
			}
		} else {

			if !a.accessGroupChecker(obj.Message.PeerID) {
				if a.accessAdminChecker(obj.Message.FromID, []string{RoleAdmin, RoleModerator, RoleNickolauyk}) {
					if msg == "/settings" {
						fmt.Println("/settings")
						err := a.sendMsgBuilder(&obj, "/settings")
						if err != nil {
							return
						}
					}
				}

				if a.accessAdminChecker(obj.Message.FromID, []string{RoleAdmin, RoleModerator, RoleNickolauyk}) {
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
				}

				if a.accessAdminChecker(obj.Message.FromID, []string{RoleAdmin, RoleModerator, RoleNickolauyk}) {
					if splitMsgs[0] == "/dialogs" {
						err := a.groupRouter(splitMsgs, obj)
						if err != nil {
							log.Error().Err(err).Msg("/dialogs")
							err = a.sendMsgBuilder(&obj, err.Error())
							if err != nil {
								return
							}
						}
					}
				}

				if a.accessAdminChecker(obj.Message.FromID, []string{RoleAdmin, RoleModerator, RoleNickolauyk}) {
					if splitMsgs[0] == "/help" {
						err := a.sendMsgBuilder(&obj, "/group info для управления группами\n/user info для управления юзерами , спойлер - для админа")
						if err != nil {
							return
						}
					}
				}

				if a.accessAdminChecker(obj.Message.FromID, []string{RoleAdmin}) {
					if splitMsgs[0] == "/user" {
						err := a.userRouter(splitMsgs, obj)
						if err != nil {
							log.Error().Err(err).Msg("/user")
							err = a.sendMsgBuilder(&obj, err.Error())
							if err != nil {
								return
							}
						}
					}
				}
			}

			if a.accessGroupChecker(obj.Message.PeerID) {
				if msg == "/Дубина" {
					b := params.NewMessagesSendBuilder()

					b.Message("@anthony_club (Антон) , вас вызывают...")
					b.RandomID(0)
					b.PeerID(obj.Message.PeerID)
					b.Attachment("photo-216838391_457239033")

					_, err := a.vk.MessagesSend(b.Params)
					if err != nil {
					}
				}
			}

			if a.accessGroupChecker(obj.Message.PeerID) {
				if msg == "/bot" {
					err := a.botHandler(obj)
					if err != nil {
						log.Error().Err(err)
						err = a.sendMsgBuilder(&obj, err.Error())
						if err != nil {
							return
						}
					}
				}
			}
		}
	})

	a.lp.MessageEvent(func(_ context.Context, obj events.MessageEventObject) {
		log.Printf("userid=%d eventid=%s payload=%s peerid=%d conversationmessageid=%d", obj.UserID, obj.EventID, obj.Payload, obj.PeerID, obj.ConversationMessageID)

		if obj.UserID == 58778743 {
			b := params.NewMessagesSendBuilder()

			b.Message("@anthony_club (Антон)")
			b.RandomID(0)
			b.PeerID(obj.PeerID)
			b.Attachment("photo-216838391_457239034")

			_, err := a.vk.MessagesSend(b.Params)
			if err != nil {
			}
		} else {

			var e MessageEventResponse
			err := json.Unmarshal(obj.Payload, &e)
			if err != nil {
				log.Error().Err(err).Msg("handler Unmarshal")
				err = a.sendMsgEventBuilder(&obj, err.Error())
				if err != nil {
					return
				}
			}

			if a.accessGroupChecker(obj.PeerID) {
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

				if e.Command == "all" {
					err = a.allHandler(obj, e.Arg)
					if err != nil {
						log.Error().Err(err).Msg("all")
						err = a.sendMsgEventBuilder(&obj, err.Error())
						if err != nil {
							return
						}
					}
				}
			}
		}
	})
}

func (a *App) sendMsgBuilder(obj *events.MessageNewObject, msg string) error {
	b := params.NewMessagesSendBuilder()

	b.Message(msg)
	//
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
