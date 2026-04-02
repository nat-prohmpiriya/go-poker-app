package service

import (
	"poker-app/model"
	"testing"
)

func TestDetermineWinnersDifferentRank(t *testing.T) {
	s := NewGameService(NewHandService())

	players := []model.Player{
		{Name: "Player 1", Hand: model.Hand{Rank: 6, Name: "Flush", TieBreak: []int{14, 10, 5, 3, 2}}},
		{Name: "Player 2", Hand: model.Hand{Rank: 2, Name: "One Pair", TieBreak: []int{11, 14, 8, 4}}},
		{Name: "Player 3", Hand: model.Hand{Rank: 1, Name: "High Card", TieBreak: []int{14, 12, 9, 5, 2}}},
		{Name: "Player 4", Hand: model.Hand{Rank: 7, Name: "Full House", TieBreak: []int{13, 10}}},
	}

	winners := s.DetermineWinners(players)

	if len(winners) != 1 || winners[0].Name != "Player 4" {
		t.Errorf("expected Player 4, got %v", winners)
	}
}

func TestDetermineWinnersTieBreak(t *testing.T) {
	s := NewGameService(NewHandService())

	// both One Pair but Player 1 has higher pair (K vs J)
	players := []model.Player{
		{Name: "Player 1", Hand: model.Hand{Rank: 2, Name: "One Pair", TieBreak: []int{13, 14, 9, 5}}},
		{Name: "Player 2", Hand: model.Hand{Rank: 2, Name: "One Pair", TieBreak: []int{11, 14, 8, 4}}},
	}

	winners := s.DetermineWinners(players)

	if len(winners) != 1 || winners[0].Name != "Player 1" {
		t.Errorf("expected Player 1 (higher pair), got %v", winners)
	}
}

func TestDetermineWinnersTieBreakKicker(t *testing.T) {
	s := NewGameService(NewHandService())

	// same pair (K) but Player 2 has higher kicker (A vs Q)
	players := []model.Player{
		{Name: "Player 1", Hand: model.Hand{Rank: 2, Name: "One Pair", TieBreak: []int{13, 12, 9, 5}}},
		{Name: "Player 2", Hand: model.Hand{Rank: 2, Name: "One Pair", TieBreak: []int{13, 14, 8, 4}}},
	}

	winners := s.DetermineWinners(players)

	if len(winners) != 1 || winners[0].Name != "Player 2" {
		t.Errorf("expected Player 2 (higher kicker), got %v", winners)
	}
}

func TestDetermineWinnersTie(t *testing.T) {
	s := NewGameService(NewHandService())

	// exact same tiebreak = tie
	players := []model.Player{
		{Name: "Player 1", Hand: model.Hand{Rank: 5, Name: "Straight", TieBreak: []int{10}}},
		{Name: "Player 2", Hand: model.Hand{Rank: 5, Name: "Straight", TieBreak: []int{10}}},
	}

	winners := s.DetermineWinners(players)

	if len(winners) != 2 {
		t.Errorf("expected 2 winners (tie), got %d", len(winners))
	}
}
