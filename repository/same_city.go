package repository

import "github.com/techoc/public_comment/model"

type SameCityRepo struct {
	DB model.DataBase
}

func (s *SameCityRepo) ListSameCity() []model.SameCity {
	var sameCities []model.SameCity

	model.DB.MyDB.Find(sameCities)
	return sameCities
}
