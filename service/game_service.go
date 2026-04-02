package service

import "poker-app/model"

type GameService struct {
	handService *HandService
}

func NewGameService(handService *HandService) *GameService {
	return &GameService{handService: handService}
}

// EvaluateAllHands evaluates hand for each player
func (s *GameService) EvaluateAllHands(players []model.Player) []model.Player {
	for i := range players {
		// extract ranks and suits from cards
		var ranks []int
		var suits []string
		for _, card := range players[i].Cards {
			ranks = append(ranks, card.Rank)
			suits = append(suits, card.Suit)
		}
		// evaluate and set hand result
		rank, name, tieBreak := s.handService.Evaluate(ranks, suits)
		players[i].Hand = model.Hand{
			Rank:     rank,
			Name:     name,
			TieBreak: tieBreak,
		}
	}
	return players
}

// DetermineWinners compares all players and returns winner(s)
// if tied on everything, returns multiple winners
func (s *GameService) DetermineWinners(players []model.Player) []model.Player {
	var winners []model.Player
	winners = append(winners, players[0])

	for i := 1; i < len(players); i++ {
		result := compareHands(winners[0], players[i])
		if result < 0 {
			// current player is stronger -> new winner
			winners = []model.Player{players[i]}
		} else if result == 0 {
			// tie -> add to winners
			winners = append(winners, players[i])
		}
		// result > 0 means current winner is still stronger
	}
	return winners
}

// compareHands compares two players' hands
// returns: -1 if b wins, 0 if tie, 1 if a wins
func compareHands(a, b model.Player) int {
	// compare hand rank first (e.g. Flush vs One Pair)
	if a.Hand.Rank > b.Hand.Rank {
		return 1
	}
	if a.Hand.Rank < b.Hand.Rank {
		return -1
	}
	// same hand rank -> compare tiebreak values one by one
	for i := 0; i < len(a.Hand.TieBreak) && i < len(b.Hand.TieBreak); i++ {
		if a.Hand.TieBreak[i] > b.Hand.TieBreak[i] {
			return 1
		}
		if a.Hand.TieBreak[i] < b.Hand.TieBreak[i] {
			return -1
		}
	}
	// all equal -> tie
	return 0
}
