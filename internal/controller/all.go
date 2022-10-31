package controller

import (
	"context"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
)

func (a *App) allHandler(obj events.MessageEventObject, group string) error {
	b := params.NewMessagesSendBuilder()
	e := params.NewMessagesSendMessageEventAnswerBuilder()

	str, err := a.logic.GetGroupForAll(context.Background(), group, int64(obj.PeerID))
	if err != nil {
		return err
	}

	b.Message(fmt.Sprintf("%s, вас вызывают: %s", group, str))
	b.RandomID(0)
	b.PeerID(obj.PeerID)

	e.EventID(obj.EventID)
	e.PeerID(obj.PeerID)
	e.UserID(obj.UserID)
	e.EventData(fmt.Sprintf(`{
    "type": "show_snackbar",
    "text": "Вы тегнули %s"
  }`, group))

	_, err = a.vk.MessagesSend(b.Params)
	if err != nil {
		return err
	}
	_, err = a.vk.MessagesSendMessageEventAnswer(e.Params)
	if err != nil {
		return err
	}

	return nil
}
