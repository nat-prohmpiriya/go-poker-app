package handler

import (
	"fmt"
	"strings"

	"poker-app/repository"
	"poker-app/service"
)

type GameHandler struct {
	deckService *service.DeckService
	gameService *service.GameService
	repo        *repository.GameRepository
}

func NewGameHandler(
	deckService *service.DeckService,
	gameService *service.GameService,
	repo *repository.GameRepository,
) *GameHandler {
	return &GameHandler{
		deckService: deckService,
		gameService: gameService,
		repo:        repo,
	}
}

// Run executes the full game flow
func (h *GameHandler) Run() {
	// 1. create deck & shuffle
	h.deckService.CreateDeck()
	h.deckService.Shuffle()

	// 2. deal 4 players x 5 cards
	h.deckService.Deal(4, 5)

	// 3. evaluate all hands
	players := h.repo.GetPlayers()
	players = h.gameService.EvaluateAllHands(players)
	h.repo.SavePlayers(players)

	// 4. print each player's hand
	for _, p := range players {
		var cards []string
		for _, c := range p.Cards {
			cards = append(cards, c.String())
		}
		fmt.Printf("%s: %s -> %s\n", p.Name, strings.Join(cards, " "), p.Hand.Name)
	}

	// 5. print cards left
	deck := h.repo.GetDeck()
	fmt.Printf("\nCards left in deck: %d\n", len(deck.Cards))

	// 6. determine and print winner(s)
	winners := h.gameService.DetermineWinners(players)
	if len(winners) == 1 {
		fmt.Printf("\n*** Winner is %s with %s! ***\n", winners[0].Name, winners[0].Hand.Name)
	} else {
		var names []string
		for _, w := range winners {
			names = append(names, w.Name)
		}
		fmt.Printf("\n*** Winners are %s with %s! ***\n", strings.Join(names, ", "), winners[0].Hand.Name)
	}
}
