package dp_service

import (
	"github.com/techoc/public_comment/model"
	"github.com/techoc/public_comment/repository"
)

type MeService struct {
	Repo *repository.MeItemRepo
}

func (m *MeService) ListMe() []model.MeItem {
	return m.Repo.ListMe()
}
