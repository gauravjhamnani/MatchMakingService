package main


type Game struct {
	id string
	maxRating int
	minCount int // minimum count in case of Multi game
	maxCount int
}

func NewGame(id string, maxRating, minCount, maxCount int) *Game {
	var game Game
	game.id = id;
	game.maxRating = maxRating
	game.minCount = minCount
	game.maxCount = maxCount
	return &game
}