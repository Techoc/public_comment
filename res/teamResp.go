package res

import "github.com/techoc/public_comment/model"

type TeamResp struct {
	model.TeamDetail
	FoodList    []TeamChooseFoodResp
	ChooseItems []TeamChooseItemResp
}
