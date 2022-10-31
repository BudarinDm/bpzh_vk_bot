package controller

import (
	"bpzh_vk_bot/internal/model/domain"
	"context"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/pkg/errors"
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
			ChatId: 0,
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
		if splitMsgs[4] != ColorBlue && splitMsgs[4] != ColorWhite && splitMsgs[4] != ColorRed && splitMsgs[4] != ColorGreen {
			return errors.New("Выбран не корректный цвет. Для подробной информации по работе с /group\nВведите /group info")
		}

		if splitMsgs[3] == "color" {
			err := a.logic.UpdateGroup(context.Background(), splitMsgs[3], splitMsgs[2], splitMsgs[4])
			if err != nil {
				return err
			}
		}

		err := a.sendMsgBuilder(&obj, fmt.Sprintf(`Обновлена группа- "%s"`, splitMsgs[2]))
		if err != nil {
			return err
		}
		return nil
	}
	if splitMsgs[1] == "delete" {
		if len(splitMsgs) < 3 {
			return errors.New("Для подробной информации по работе с /group\nВведите /group info")
		}
		err := a.logic.DeleteGroup(context.Background(), splitMsgs[2])
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
		if len(splitMsgs) < 3 {
			return errors.New("Для подробной информации по работе с /group\nВведите /group info")
		}

	}
	if splitMsgs[1] == "list" {
		listMsg :=
			"Доступные беседы:\nname: Тестовая беседа, id: 2000000002\nname: Поддержка Бота, id: 2000000001"
		err := a.sendMsgBuilder(&obj, listMsg)
		if err != nil {
			return err
		}
		return nil
	}
	if splitMsgs[1] == "info" {
		infoMsg :=
			"Доступные команды для group:\ncreate- /group create [айди группы или название из списка /group list] [название группы] [цвет группы из доступных],\n" +
				"update- /group update [название группы] [color] [новое значение],\n" +
				"delete- /group delete [название группы].\n" +
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
