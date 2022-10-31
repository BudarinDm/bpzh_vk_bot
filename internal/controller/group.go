package controller

import (
	"bpzh_vk_bot/internal/model/domain"
	"context"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/pkg/errors"
	"strconv"
)

func (a *App) groupRouter(splitMsgs []string, obj events.MessageNewObject) error {
	if len(splitMsgs) < 2 {
		return errors.New("Для подробной информации по работе с /group\nВведите /group info")
	}

	if splitMsgs[1] == "create" {
		if len(splitMsgs) < 4 {
			return errors.New("Для подробной информации по работе с /group\nВведите /group info")
		}
		if splitMsgs[3] != ColorBlue && splitMsgs[3] != ColorWhite && splitMsgs[3] != ColorRed && splitMsgs[3] != ColorGreen {
			return errors.New("Выбран не корректный цвет. Для подробной информации по работе с /group\nВведите /group info")
		}

		g := domain.Group{
			Name:   splitMsgs[2],
			Color:  splitMsgs[3],
			ChatId: int64(obj.Message.PeerID),
		}

		item, err := a.logic.CreateGroup(context.Background(), g)
		if err != nil {
			return err
		}
		var msg string
		if item == nil {
			msg = fmt.Sprintf(`Создана группа -"%s"", цвет группы -"%s"`, splitMsgs[2], splitMsgs[3])
		} else {
			msg = fmt.Sprintf(`Такая группа уже существует- "%s", цвет группы- "%s"`, item.Name, item.Color)
		}

		err = a.sendMsgBuilder(&obj, msg)
		if err != nil {
			return err
		}

		return nil
	}
	if splitMsgs[1] == "update" {
		if len(splitMsgs) < 5 {
			return errors.New("Для подробной информации по работе с /group\nВведите /group info")
		}

		if splitMsgs[3] != "color" && splitMsgs[3] != "name" {
			return errors.New("Можно обновить или название или цвет /group\nВведите /group info")

		}
		if splitMsgs[3] == "color" {
			if splitMsgs[4] != ColorBlue && splitMsgs[4] != ColorWhite && splitMsgs[4] != ColorRed && splitMsgs[4] != ColorGreen {
				return errors.New("Выбран не корректный цвет. Для подробной информации по работе с /group\nВведите /group info")
			}
		}
		err := a.logic.UpdateGroup(context.Background(), splitMsgs[3], splitMsgs[2], splitMsgs[4], int64(obj.Message.PeerID))
		if err != nil {
			return err
		}

		err = a.sendMsgBuilder(&obj, fmt.Sprintf(`Обновлена группа- "%s"`, splitMsgs[2]))
		if err != nil {
			return err
		}
		return nil
	}
	if splitMsgs[1] == "delete" {
		if len(splitMsgs) < 3 {
			return errors.New("Для подробной информации по работе с /group\nВведите /group info")
		}
		err := a.logic.DeleteGroup(context.Background(), splitMsgs[2], int64(obj.Message.PeerID))
		if err != nil {
			return err
		}

		err = a.sendMsgBuilder(&obj, fmt.Sprintf(`Удалена группа- "%s"`, splitMsgs[2]))
		if err != nil {
			return err
		}
		return nil
	}
	if splitMsgs[1] == "add" {
		if len(splitMsgs) < 4 {
			return errors.New("Для подробной информации по работе с /group\nВведите /group info")
		}

		_, err := a.logic.GetGroup(context.Background(), splitMsgs[3], int64(obj.Message.PeerID))
		if err != nil {
			return errors.New(fmt.Sprintf(`Группа "%s" не найдена`, splitMsgs[3]))
		}

		id, err := strconv.ParseInt(splitMsgs[2], 10, 64)
		if err != nil {
			return err
		}

		err = a.logic.AddUserToGroup(context.Background(), id, splitMsgs[3], int64(obj.Message.PeerID))
		if err != nil {
			return err
		}
		err = a.sendMsgBuilder(&obj, fmt.Sprintf("Пользователь %d добавлен в группу %s", id, splitMsgs[3]))
		if err != nil {
			return err
		}
		return nil
	}
	if splitMsgs[1] == "kick" {
		if len(splitMsgs) < 4 {
			return errors.New("Для подробной информации по работе с /group\nВведите /group info")
		}

		_, err := a.logic.GetGroup(context.Background(), splitMsgs[3], int64(obj.Message.PeerID))
		if err != nil {
			return errors.New(fmt.Sprintf(`Группа "%s" не найдена`, splitMsgs[3]))
		}

		id, err := strconv.ParseInt(splitMsgs[2], 10, 64)
		if err != nil {
			return err
		}

		err = a.logic.DeleteUserToGroup(context.Background(), id, splitMsgs[3], int64(obj.Message.PeerID))
		if err != nil {
			return err
		}
		err = a.sendMsgBuilder(&obj, fmt.Sprintf("Пользователь %d удален из группы %s", id, splitMsgs[3]))
		if err != nil {
			return err
		}
		return nil
	}
	if splitMsgs[1] == "info" {
		infoMsg :=
			"Доступные команды для group:\nСоздать группу-  /group create [название группы] [цвет группы из доступных],\n" +
				"Обновить группу-  /group update [название группы] [color] [новое значение],\n" +
				"Удалить группу-  /group delete [название группы].\n" +
				"Добавить в группу-  /group add [id юзера] [название группы]\n" +
				"Удалить из группы-  /group kick [id юзера] [название группы]\n" +
				"Доступные цвета: primary — синяя, secondary — белая, negative — красный, positive — зеленый"
		err := a.sendMsgBuilder(&obj, infoMsg)
		if err != nil {
			return err
		}
		return nil
	}

	err := a.sendMsgBuilder(&obj, "Для подробной информации по работе с /group\nВведите /group info")
	if err != nil {
		return err
	}
	return nil
}
