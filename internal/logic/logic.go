package logic

import (
	"bpzh_vk_bot/internal/config"
	"bpzh_vk_bot/internal/repo/db"
	"github.com/SevereCloud/vksdk/v2/api"
)

// Logic содержит все для доступа к данным
type Logic struct {
	config *config.Config
	repo   *db.Repo
	vk     *api.VK
}

func NewLogic(config *config.Config, repo *db.Repo, vk *api.VK) *Logic {
	return &Logic{
		config: config,
		repo:   repo,
		vk:     vk,
	}
}
