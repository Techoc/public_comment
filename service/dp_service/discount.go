package dp_service

import (
	"github.com/techoc/public_comment/model"
	"github.com/techoc/public_comment/repository"
)

type DiscountService struct {
	Repo *repository.Discount
}

func (d *DiscountService) ListDiscounts() []model.Discount {
	itemsLeft := d.Repo.ListDiscounts(1)
	itemsRight := d.Repo.ListDiscounts(2)
	itemsLeft = append(itemsLeft, itemsRight...)
	return itemsLeft
}
