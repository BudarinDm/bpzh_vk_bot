package db

import (
	"bpzh_vk_bot/internal/model/domain"
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
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

		var user domain.User
		err = doc.DataTo(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *Repo) GetAdminUser(ctx context.Context, userId int) (*domain.User, error) {
	var user domain.User
	iter := r.FS.Collection("users").Where("vkid", "==", userId).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		err = doc.DataTo(&user)
		if err != nil {
			return nil, err
		}

	}

	return &user, nil
}

func (r *Repo) AddRoleUser(ctx context.Context, userId, chatId int64, role string) error {
	iter := r.FS.Collection("users").Where("vkid", "==", userId).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		var user domain.User
		fmt.Println(doc.Data())
		err = doc.DataTo(&user)
		if err != nil {
			return err
		}

		_, err = r.FS.Collection("users").Doc(doc.Ref.ID).Set(ctx, map[string]interface{}{
			"roles": map[string]interface{}{
				role: firestore.ArrayUnion(chatId),
			},
		}, firestore.MergeAll)
		if err != nil {
			return err
		}
		return nil

	}
	return nil
}

func (r *Repo) DeleteRoleUser(ctx context.Context, userId, chatId int64, role string) error {
	iter := r.FS.Collection("users").Where("vkid", "==", userId).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		var user domain.User
		err = doc.DataTo(&user)
		if err != nil {
			return err
		}

		_, err = r.FS.Collection("users").Doc(doc.Ref.ID).Set(ctx, map[string]interface{}{
			"roles": map[string]interface{}{
				role: firestore.ArrayRemove(chatId),
			},
		}, firestore.MergeAll)
		if err != nil {
			return err
		}
		return nil

	}
	return nil
}

func (r *Repo) CreateUser(ctx context.Context, u domain.User) error {
	_, _, err := r.FS.Collection("users").Add(ctx, u)
	if err != nil {
		return err
	}

	return nil
}
