package db

import (
	"bpzh_vk_bot/internal/model/domain"
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *Repo) GetGroups(ctx context.Context) ([]domain.Group, error) {
	var items []domain.Group
	iter := r.FS.Collection("groups").Documents(ctx)
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
	dsnap, err := r.FS.Collection("groups").Doc(g.Name).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			_, err = r.FS.Collection("groups").Doc(g.Name).Set(ctx, g)
			if err != nil {
				return nil, err
			}
			return nil, nil
		}
	}
	data := dsnap.Data()
	var item domain.Group
	ln, ok := data["name"]
	if ok {
		item.Name = ln.(string)
	}
	c, ok := data["color"]
	if ok {
		item.Color = c.(string)
	}

	return &item, nil
}

func (r *Repo) DeleteGroup(ctx context.Context, g string) error {
	_, err := r.FS.Collection("groups").Doc(g).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateGroup(ctx context.Context, field, nameGroup, value string) error {
	if field == "color" {
		_, err := r.FS.Collection("groups").Doc(nameGroup).Update(ctx, []firestore.Update{
			{
				Path:  field,
				Value: value,
			},
		})
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (r *Repo) AddUserToGroup(ctx context.Context, userId int64, group string) error {
	_, err := r.FS.Collection("groups").Doc(group).Set(ctx, map[string]interface{}{
		"users": firestore.ArrayUnion(userId),
	}, firestore.MergeAll)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetGroup(ctx context.Context, name string) (*domain.Group, error) {
	var gr domain.Group
	iter := r.FS.Collection("groups").Where(name, "==", name).Documents(ctx)
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
