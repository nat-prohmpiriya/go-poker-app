package service

import (
	"poker-app/repository"
	"testing"
)

func TestCreateDeck52Cards(t *testing.T) {
	repo := repository.NewGameRepository()
	s := NewDeckService(repo)

	s.CreateDeck()
	deck := repo.GetDeck()

	if len(deck.Cards) != 52 {
		t.Errorf("expected 52 cards, got %d", len(deck.Cards))
	}
}

func TestCreateDeckNoDuplicates(t *testing.T) {
	repo := repository.NewGameRepository()
	s := NewDeckService(repo)

	s.CreateDeck()
	deck := repo.GetDeck()

	seen := make(map[string]bool)
	for _, card := range deck.Cards {
		key := card.String()
		if seen[key] {
			t.Errorf("duplicate card found: %s", key)
		}
		seen[key] = true
	}
}

func TestShuffleChangeOrder(t *testing.T) {
	repo := repository.NewGameRepository()
	s := NewDeckService(repo)

	s.CreateDeck()
	before := make([]string, 52)
	for i, c := range repo.GetDeck().Cards {
		before[i] = c.String()
	}

	s.Shuffle()
	after := repo.GetDeck().Cards

	same := 0
	for i, c := range after {
		if c.String() == before[i] {
			same++
		}
	}

	// after shuffle, most cards should be in different positions
	if same > 10 {
		t.Errorf("shuffle did not change enough card positions (%d/52 same)", same)
	}
}

func TestDeal4Players5Cards(t *testing.T) {
	repo := repository.NewGameRepository()
	s := NewDeckService(repo)

	s.CreateDeck()
	s.Shuffle()
	s.Deal(4, 5)

	players := repo.GetPlayers()
	deck := repo.GetDeck()

	if len(players) != 4 {
		t.Errorf("expected 4 players, got %d", len(players))
	}
	for i, p := range players {
		if len(p.Cards) != 5 {
			t.Errorf("player %d should have 5 cards, got %d", i+1, len(p.Cards))
		}
	}
	if len(deck.Cards) != 32 {
		t.Errorf("expected 32 cards left, got %d", len(deck.Cards))
	}
}

func TestDrawOne(t *testing.T) {
	repo := repository.NewGameRepository()
	s := NewDeckService(repo)

	s.CreateDeck()
	s.Shuffle()
	s.InitPlayers(2)

	card := s.DrawOne(0)
	players := repo.GetPlayers()
	deck := repo.GetDeck()

	if len(players[0].Cards) != 1 {
		t.Errorf("player 1 should have 1 card, got %d", len(players[0].Cards))
	}
	if players[0].Cards[0] != card {
		t.Errorf("drawn card should match player's card")
	}
	if len(deck.Cards) != 51 {
		t.Errorf("expected 51 cards left, got %d", len(deck.Cards))
	}
}
