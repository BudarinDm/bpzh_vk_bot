package logic

import (
	"bpzh_vk_bot/internal/model/domain"
	"context"
)

func (l *Logic) GetGroups(ctx context.Context) ([]domain.Group, error) {
	return l.repo.GetGroups(ctx)
}

func (l *Logic) CreateGroup(ctx context.Context, g domain.Group) (*domain.Group, error) {
	return l.repo.CreateGroup(ctx, g)
}

func (l *Logic) DeleteGroup(ctx context.Context, g string) error {
	return l.repo.DeleteGroup(ctx, g)
}
func (l *Logic) UpdateGroup(ctx context.Context, field, nameGroup, value string) error {
	return l.repo.UpdateGroup(ctx, field, nameGroup, value)
}

func (l *Logic) AddUserToGroup(ctx context.Context, userId int64, group string) error {
	return l.repo.AddUserToGroup(ctx, userId, group)
}

func (l *Logic) GetGroup(ctx context.Context, name string) (*domain.Group, error) {
	return l.repo.GetGroup(ctx, name)
}

func (l *Logic) DeleteUserToGroup(ctx context.Context, userId int64, group string) error {
	return l.repo.DeleteUserToGroup(ctx, userId, group)
}
