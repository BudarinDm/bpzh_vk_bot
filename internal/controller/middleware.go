package controller

import "context"

func (a *App) accessMethodGroupChecker(objId int) bool {
	for _, cid := range AccessChatIds {
		if cid == int64(objId) {
			return true
		}
	}
	return false
}

func (a *App) accessMethodAdminChecker(userId int, adminGroupName string) bool {
	users, err := a.logic.GetAdminUsers(context.Background(), adminGroupName)
	if err != nil {
		return false
	}

	for _, u := range users {
		if u.VkId == int64(userId) {
			return true
		}
	}
	return false
}
