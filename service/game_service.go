package service

import (
	"poker-app/model"
	"poker-app/repository"
)

type GameService struct {
	handService *HandService
	repo        *repository.GameRepository
}

func NewGameService(handService *HandService, repo *repository.GameRepository) *GameService {
	return &GameService{handService: handService, repo: repo}
}

func (s *GameService) GetPlayers() []model.Player {
	return s.repo.GetPlayers()
}

func (s *GameService) SavePlayers(players []model.Player) {
	s.repo.SavePlayers(players)
}

func (s *GameService) EvaluateAllHands(players []model.Player) []model.Player {
	for i := range players {
		var ranks []int
		var suits []string
		for _, card := range players[i].Cards {
			ranks = append(ranks, card.Rank)
			suits = append(suits, card.Suit)
		}
		rank, name, tieBreak := s.handService.Evaluate(ranks, suits)
		players[i].Hand = model.Hand{
			Rank:     rank,
			Name:     name,
			TieBreak: tieBreak,
		}
	}
	return players
}

func (s *GameService) DetermineWinners(players []model.Player) []model.Player {
	var winners []model.Player
	winners = append(winners, players[0])

	for i := 1; i < len(players); i++ {
		result := compareHands(winners[0], players[i])
		if result < 0 {
			winners = []model.Player{players[i]}
		} else if result == 0 {
			winners = append(winners, players[i])
		}
	}
	return winners
}

// returns: -1 if b wins, 0 if tie, 1 if a wins
func compareHands(a, b model.Player) int {
	if a.Hand.Rank > b.Hand.Rank {
		return 1
	}
	if a.Hand.Rank < b.Hand.Rank {
		return -1
	}
	// same hand rank -> compare tiebreak
	for i := 0; i < len(a.Hand.TieBreak) && i < len(b.Hand.TieBreak); i++ {
		if a.Hand.TieBreak[i] > b.Hand.TieBreak[i] {
			return 1
		}
		if a.Hand.TieBreak[i] < b.Hand.TieBreak[i] {
			return -1
		}
	}
	return 0
}
