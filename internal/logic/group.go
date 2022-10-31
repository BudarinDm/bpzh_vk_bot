package logic

import (
	"bpzh_vk_bot/internal/model/domain"
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
)

func (l *Logic) GetGroupsByChatId(ctx context.Context, chatId int64) ([]domain.Group, error) {
	return l.repo.GetGroupsByChatId(ctx, chatId)
}

func (l *Logic) CreateGroup(ctx context.Context, g domain.Group) (*domain.Group, error) {
	return l.repo.CreateGroup(ctx, g)
}

func (l *Logic) DeleteGroup(ctx context.Context, g string, chatId int64) error {
	return l.repo.DeleteGroup(ctx, g, chatId)
}
func (l *Logic) UpdateGroup(ctx context.Context, field, nameGroup, value string, chatId int64) error {
	return l.repo.UpdateGroup(ctx, field, nameGroup, value, chatId)
}

func (l *Logic) AddUserToGroup(ctx context.Context, userId int64, group string, chatId int64) error {
	return l.repo.AddUserToGroup(ctx, userId, group, chatId)
}

func (l *Logic) GetGroup(ctx context.Context, name string, chatId int64) (*domain.Group, error) {
	return l.repo.GetGroup(ctx, name, chatId)
}

func (l *Logic) DeleteUserToGroup(ctx context.Context, userId int64, group string, chatId int64) error {
	return l.repo.DeleteUserToGroup(ctx, userId, group, chatId)
}

func (l *Logic) GetGroupForAll(ctx context.Context, name string, chatId int64) (string, error) {
	var resp string

	group, err := l.repo.GetGroup(context.Background(), name, chatId)
	if err != nil {
		return "", err
	}

	for _, g := range group.Users {
		user, err := l.repo.GetUserVKDomain(ctx, g)
		if err != nil {
			log.Error().Err(err)
		}
		resp += fmt.Sprintf("@%s (%s) ", user.VkDomain, user.FirstName)
	}

	return resp, nil
}
