package repository

import (
	"github.com/techoc/public_comment/model"
)

type MarketRepo struct {
	DB model.DataBase
}

func (m *MarketRepo) GetMarketInfo(hotelId string) model.Market {
	var market model.Market
	m.DB.MyDB.Where("hotel_id=?", hotelId).Find(&market)
	return market
}
