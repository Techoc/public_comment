package repository

import "github.com/techoc/public_comment/model"

type IndexNavItemRepo struct {
	DB model.DataBase
}

func (item *IndexNavItemRepo) ListIndexNavItem() ([]*model.IndexNavItem, uint64, error) {
	NavItems := make([]*model.IndexNavItem, 0)
	var count uint64

	if err := model.DB.MyDB.Model(NavItems).Count(count).Error; err != nil {
		return nil, 0, err
	}
	return NavItems, count, nil
}
