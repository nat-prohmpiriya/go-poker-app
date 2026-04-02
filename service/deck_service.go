package service

import (
	"fmt"
	"math/rand"
	"poker-app/model"
	"poker-app/repository"
)

type DeckService struct {
	repo *repository.GameRepository
}

func NewDeckService(repo *repository.GameRepository) *DeckService {
	return &DeckService{repo: repo}
}

func (s *DeckService) CreateDeck() {
	var cards []model.Card
	for _, suit := range model.Suits {
		for rank := 2; rank <= 14; rank++ {
			cards = append(cards, model.Card{Rank: rank, Suit: suit})
		}
	}
	s.repo.SaveDeck(model.Deck{Cards: cards})
}

func (s *DeckService) Shuffle() {
	deck := s.repo.GetDeck()
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
	s.repo.SaveDeck(deck)
}

func (s *DeckService) Deal(numPlayers, cardsPerPlayer int) {
	deck := s.repo.GetDeck()
	var players []model.Player

	for i := 0; i < numPlayers; i++ {
		hand := deck.Cards[:cardsPerPlayer]
		deck.Cards = deck.Cards[cardsPerPlayer:]
		players = append(players, model.Player{
			Name:  fmt.Sprintf("Player %d", i+1),
			Cards: hand,
		})
	}
	s.repo.SaveDeck(deck)
	s.repo.SavePlayers(players)
}

func (s *DeckService) InitPlayers(numPlayers int) {
	var players []model.Player
	for i := 0; i < numPlayers; i++ {
		players = append(players, model.Player{
			Name: fmt.Sprintf("Player %d", i+1),
		})
	}
	s.repo.SavePlayers(players)
}

func (s *DeckService) DrawOne(playerIndex int) model.Card {
	deck := s.repo.GetDeck()
	players := s.repo.GetPlayers()

	card := deck.Cards[0]
	deck.Cards = deck.Cards[1:]
	players[playerIndex].Cards = append(players[playerIndex].Cards, card)

	s.repo.SaveDeck(deck)
	s.repo.SavePlayers(players)

	return card
}
