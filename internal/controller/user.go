package controller

import (
	"context"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/pkg/errors"
	"strconv"
)

func (a *App) userRouter(splitMsgs []string, obj events.MessageNewObject) error {
	if len(splitMsgs) < 2 {
		return errors.New("Для подробной информации по работе с /user\nВведите /user info")
	}

	if splitMsgs[1] == "role" {
		if len(splitMsgs) < 5 {
			return errors.New("Для подробной информации по работе с /user\nВведите /user info")
		}
		if splitMsgs[2] == "add" {
			if splitMsgs[4] != RoleAdmin && splitMsgs[4] != RoleModerator && splitMsgs[4] != RoleNickolauyk {
				return errors.New("Выбрана не корректная роль. Для подробной информации по работе с /user\nВведите /user info")
			}

			userId, err := strconv.ParseInt(splitMsgs[3], 10, 64)
			if err != nil {
				return errors.New("Передан не корректный айди")
			}
			err = a.logic.AddRoleUser(context.Background(), userId, int64(obj.Message.PeerID), splitMsgs[4])
			if err != nil {
				return err
			}

			msg := fmt.Sprintf(`Пользователю %d, добавленна роль %s`, userId, splitMsgs[4])

			err = a.sendMsgBuilder(&obj, msg)
			if err != nil {
				return err
			}

			return nil
		}
		if splitMsgs[2] == "delete" {
			if splitMsgs[4] != RoleAdmin && splitMsgs[4] != RoleModerator && splitMsgs[4] != RoleNickolauyk {
				return errors.New("Выбрана не корректная роль. Для подробной информации по работе с /user\nВведите /user info")
			}

			userId, err := strconv.ParseInt(splitMsgs[3], 10, 64)
			if err != nil {
				return errors.New("Передан не корректный айди")
			}
			err = a.logic.DeleteRoleUser(context.Background(), userId, int64(obj.Message.PeerID), splitMsgs[4])
			if err != nil {
				return err
			}

			msg := fmt.Sprintf(`У пользователя %d, удалена роль %s`, userId, splitMsgs[4])

			err = a.sendMsgBuilder(&obj, msg)
			if err != nil {
				return err
			}

			return nil
		}
	}
	if splitMsgs[1] == "info" {
		infoMsg :=
			"Доступные команды для user:\nДобавить пользователю роль- /user role add [айди пользователя] [название роли]\n" +
				"Удалить роль у пользователя- /user role delete [айди пользователя] [название роли].\n" +
				"Доступные роли- /user info roles\n" +
				"Список пользователей- /user info users"
		if len(splitMsgs) == 3 && splitMsgs[2] == "roles" {
			infoMsg = "Доступные роли:\nadmin, moderator, Николаюк"
		}
		if len(splitMsgs) == 3 && splitMsgs[2] == "users" {
			users, err := a.logic.GetChatUsers(obj.Message.PeerID)
			if err != nil {
				return err
			}
			infoMsg = "Пользователи: \n"
			for _, u := range users {
				infoMsg += fmt.Sprintf("%d: %s %s\n", u.VkId, u.LastName, u.FirstName)
			}
		}

		err := a.sendMsgBuilder(&obj, infoMsg)
		if err != nil {
			return err
		}
		return nil
	}

	err := a.sendMsgBuilder(&obj, "Для подробной информации по работе с /user\nВведите /user info")
	if err != nil {
		return err
	}
	return nil
}
