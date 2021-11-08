package repository

import "github.com/techoc/public_comment/model"

type GuessRepo struct {
	DB model.DataBase
}

func (g *GuessRepo) ListGuesses() []model.Guess {
	var guessList []model.Guess
	g.DB.MyDB.Find(&guessList)
	return guessList
}
