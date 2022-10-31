package controller

import "context"

func (a *App) accessAdminChecker(userId, chatId int, roles []string) bool {
	user, err := a.logic.GetAdminUser(context.Background(), userId)
	if err != nil {
		return false
	}

	for _, v := range roles {
		rls, _ := user.Roles[v]
		for _, r := range rls {
			if r == int64(chatId) {
				return true
			}
		}
	}

	return false
}
