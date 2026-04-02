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

// แจกไพ่ให้ผู้เล่น 4 คน คนละ ห้าใบ
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
