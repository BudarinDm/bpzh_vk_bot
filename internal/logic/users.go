package logic

import (
	"bpzh_vk_bot/internal/model/domain"
	"context"
	"github.com/SevereCloud/vksdk/v2/api"
	"strconv"
)

func (l *Logic) GetUsers(ctx context.Context) ([]domain.User, error) {
	return l.repo.GetUsers(ctx)
}

func (l *Logic) GetAdminUser(ctx context.Context, userId int) (*domain.User, error) {
	return l.repo.GetAdminUser(ctx, userId)
}

func (l *Logic) AddRoleUser(ctx context.Context, userId, chatId int64, role string) error {
	return l.repo.AddRoleUser(ctx, userId, chatId, role)
}

func (l *Logic) DeleteRoleUser(ctx context.Context, userId, chatId int64, role string) error {
	return l.repo.DeleteRoleUser(ctx, userId, chatId, role)
}

func (l *Logic) CreateUser(ctx context.Context, u domain.User) error {
	return l.repo.CreateUser(ctx, u)
}

func (l *Logic) GetChatUsers(chatId int) ([]domain.User, error) {
	members, err := l.vk.MessagesGetConversationMembers(api.Params{"peer_id": chatId})
	if err != nil {
		return nil, err
	}

	var usersIds string
	for _, m := range members.Items {
		if m.MemberID < 0 {
			continue
		}
		usersIds += "," + strconv.Itoa(m.MemberID)
	}
	resp, err := l.vk.UsersGet(api.Params{"user_ids": usersIds})
	if err != nil {
		return nil, err
	}
	var users []domain.User
	for _, u := range resp {
		user := domain.User{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			VkId:      int64(u.ID),
		}
		users = append(users, user)
	}

	return users, nil
}
