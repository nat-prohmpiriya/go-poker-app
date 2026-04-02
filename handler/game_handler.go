package handler

import (
	"fmt"
	"strings"

	"poker-app/model"
	"poker-app/service"
)

type GameHandler struct {
	deckService *service.DeckService
	gameService *service.GameService
}

func NewGameHandler(
	deckService *service.DeckService,
	gameService *service.GameService,
) *GameHandler {
	return &GameHandler{
		deckService: deckService,
		gameService: gameService,
	}
}

func (h *GameHandler) Run() {
	h.deckService.CreateDeck()
	h.deckService.Shuffle()
	h.deckService.Deal(4, 5)

	players := h.gameService.GetPlayers()
	players = h.gameService.EvaluateAllHands(players)
	h.gameService.SavePlayers(players)

	for _, p := range players {
		fmt.Printf("%s: %s -> %s\n", p.Name, formatCards(p.Cards), p.Hand.Name)
	}

	fmt.Printf("\nCards left in deck: %d\n", h.deckService.GetCardsLeft())

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

func formatCards(cards []model.Card) string {
	var parts []string
	for _, c := range cards {
		parts = append(parts, c.String())
	}
	return strings.Join(parts, " ")
}
