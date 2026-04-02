package repository

import "poker-app/model"

type GameRepository struct {
	Deck    model.Deck
	Players []model.Player
}

func NewGameRepository() *GameRepository {
	return &GameRepository{}
}

func (r *GameRepository) SaveDeck(deck model.Deck) {
	r.Deck = deck
}

func (r *GameRepository) GetDeck() model.Deck {
	return r.Deck
}

func (r *GameRepository) SavePlayers(players []model.Player) {
	r.Players = players
}

func (r *GameRepository) GetPlayers() []model.Player {
	return r.Players
}
