package controller

import "context"

func (a *App) accessAdminChecker(userId int, roles []string) bool {
	user, err := a.logic.GetAdminUser(context.Background(), userId)
	if err != nil {
		return false
	}

	for _, v := range roles {
		for _, r := range user.Roles {
			if v == r {
				return true
			}
		}
	}
	return false
}

func (a *App) accessGroupChecker(groupId int) bool {
	for _, c := range AccessChatIds {
		if c == groupId {
			return true
		}
	}
	return false
}
