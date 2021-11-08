package repository

import (
	"github.com/techoc/public_comment/model"
)

type ListNavItemRepo struct {
	DB model.DataBase
}

func (i *ListNavItemRepo) ListNavItems(level int) []model.Nav {
	var items []model.Nav
	i.DB.MyDB.Where("level=?", level).Find(&items)
	return items
}
