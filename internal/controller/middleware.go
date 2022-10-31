package controller

import "context"

func (a *App) accessMethodGroupChecker(objId int64) bool {
	for _, cid := range AccessChatIds {
		if cid == objId {
			return true
		}
	}
	return false
}

func (a *App) accessMethodAdminChecker(userId int64, adminGroupName string) bool {
	users, err := a.logic.GetAdminUsers(context.Background(), adminGroupName)
	if err != nil {
		return false
	}

	for _, u := range users {
		if u.VkId == userId {
			return true
		}
	}
	return false
}
