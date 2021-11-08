package dp_service

import (
	"github.com/techoc/public_comment/model"
	"github.com/techoc/public_comment/repository"
)

type GuessService struct {
	Repo *repository.GuessRepo
}

func (g *GuessService) ListGuess() []model.Guess {
	return g.Repo.ListGuesses()
}
