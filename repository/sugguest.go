package repository

import "github.com/techoc/public_comment/model"

type SuggestRepo struct {
	DB model.DataBase
}

func (s *SuggestRepo) ListSuggest(level int) model.Suggest {
	var item model.Suggest
	model.DB.MyDB.Where("level=?", level).Find(&item)
	return item
}
