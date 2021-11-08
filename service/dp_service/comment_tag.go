package dp_service

import (
	"github.com/techoc/public_comment/model"
	"github.com/techoc/public_comment/repository"
)

type CommentTagService struct {
	Repo repository.CommentTagRepo
}

func (c *CommentTagService) GetCommentTagList(hotelId string) []model.CommentTag {
	return c.Repo.GetCommentTagList(hotelId)
}
