package dp_service

import (
	"github.com/techoc/public_comment/model"
	"github.com/techoc/public_comment/repository"
)

type CommentService struct {
	Repo repository.CommentRepo
}

func (c *CommentService) GetCommentList(hotelId string) []model.Comment {
	return c.Repo.GetCommentList(hotelId)
}
