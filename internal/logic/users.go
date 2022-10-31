package logic

import (
	"bpzh_vk_bot/internal/model/domain"
	"context"
)

func (l *Logic) GetUsers(ctx context.Context) ([]domain.User, error) {
	return l.repo.GetUsers(ctx)
}
