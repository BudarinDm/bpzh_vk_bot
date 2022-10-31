package db

import (
	"bpzh_vk_bot/internal/model/domain"
	"context"
	"google.golang.org/api/iterator"
)

func (r *Repo) GetUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	iter := r.FS.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		data := doc.Data()

		var user domain.User
		id, ok := data["id"]
		if ok {
			user.Id = id.(int64)
		}
		ln, ok := data["lastname"]
		if ok {
			user.LastName = ln.(string)
		}
		fn, ok := data["firstname"]
		if ok {
			user.FirstName = fn.(string)
		}
		tgid, ok := data["tgid"]
		if ok {
			user.TgId = tgid.(string)
		}
		vkid, ok := data["vkid"]
		if ok {
			user.VkId = vkid.(string)
		}

		users = append(users, user)
	}

	return users, nil
}
