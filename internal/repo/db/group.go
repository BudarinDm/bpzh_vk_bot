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
			_, err = r.FS.Collection("groups").Doc(g.Name).Set(ctx, map[string]interface{}{
				"name":  g.Name,
				"color": g.Color,
			})
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
