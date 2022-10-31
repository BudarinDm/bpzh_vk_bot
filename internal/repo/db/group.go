package db

import (
	"bpzh_vk_bot/internal/model/domain"
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
)

func (r *Repo) GetGroupsByChatId(ctx context.Context, chatId int64) ([]domain.Group, error) {
	var items []domain.Group
	iter := r.FS.Collection("groups").Where("chatid", "==", chatId).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		data := doc.Data()

		var item domain.Group
		ln, ok := data["name"]
		if ok {
			item.Name = ln.(string)
		}
		c, ok := data["color"]
		if ok {
			item.Color = c.(string)
		}

		items = append(items, item)
	}
	return items, nil
}

func (r *Repo) CreateGroup(ctx context.Context, g domain.Group) (*domain.Group, error) {
	iter := r.FS.Collection("groups").Where("name", "==", g.Name).Where("chatid", "==", g.ChatId).Documents(ctx)

	var groups []domain.Group
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var gr domain.Group
		err = doc.DataTo(&gr)
		if err != nil {
			return nil, err
		}

		groups = append(groups, gr)
	}

	if len(groups) == 0 {
		_, _, err := r.FS.Collection("groups").Add(ctx, g)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	return &groups[0], nil
}

func (r *Repo) DeleteGroup(ctx context.Context, g string, chatId int64) error {
	iter := r.FS.Collection("groups").Where("name", "==", g).Where("chatid", "==", chatId).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}
		_, err = r.FS.Collection("groups").Doc(doc.Ref.ID).Delete(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repo) UpdateGroup(ctx context.Context, field, nameGroup, value string, chatId int64) error {
	iter := r.FS.Collection("groups").Where("name", "==", nameGroup).Where("chatid", "==", chatId).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}
		_, err = r.FS.Collection("groups").Doc(doc.Ref.ID).Update(ctx, []firestore.Update{
			{
				Path:  field,
				Value: value,
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repo) AddUserToGroup(ctx context.Context, userId int64, group string, chatId int64) error {
	iter := r.FS.Collection("groups").Where("name", "==", group).Where("chatid", "==", chatId).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		if err != nil {
			return err
		}
		_, err = r.FS.Collection("groups").Doc(doc.Ref.ID).Set(ctx, map[string]interface{}{
			"users": firestore.ArrayUnion(userId),
		}, firestore.MergeAll)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repo) GetGroup(ctx context.Context, name string, chatId int64) (*domain.Group, error) {
	var gr domain.Group
	iter := r.FS.Collection("groups").Where("name", "==", name).Where("chatid", "==", chatId).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		err = doc.DataTo(&gr)
		if err != nil {
			return nil, err
		}
	}

	return &gr, nil
}

func (r *Repo) DeleteUserToGroup(ctx context.Context, userId int64, group string, chatId int64) error {
	iter := r.FS.Collection("groups").Where("name", "==", group).Where("chatid", "==", chatId).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		_, err = r.FS.Collection("groups").Doc(doc.Ref.ID).Set(ctx, map[string]interface{}{
			"users": firestore.ArrayRemove(userId),
		}, firestore.MergeAll)
		if err != nil {
			return err
		}
	}
	return nil
}
