package handler

import (
	"bufio"
	"fmt"
	"os"
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
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== POKER GAME ===")
		fmt.Println("[1] Start New Round")
		fmt.Println("[2] Exit")
		fmt.Print("\n> ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			h.playRound(scanner)
		case "2":
			fmt.Println("\nThanks for playing!")
			return
		default:
			fmt.Println("\nInvalid choice. Please enter 1 or 2.")
		}
	}
}

func (h *GameHandler) playRound(scanner *bufio.Scanner) {
	fmt.Println("\nShuffling deck...")
	h.deckService.CreateDeck()
	h.deckService.Shuffle()

	fmt.Printf("Cards in deck: %d\n", h.deckService.GetCardsLeft())

	h.deckService.InitPlayers(4)
	cardsPerPlayer := 5

	// Player 1 (You) draws interactively
	fmt.Println("\n--- Your turn (Player 1) ---")
	for j := 0; j < cardsPerPlayer; j++ {
		h.showPlayerMenu(0, j+1, cardsPerPlayer, scanner)
	}
	players := h.gameService.GetPlayers()
	fmt.Printf("\nYour hand: %s\n", formatCards(players[0].Cards))

	// Bot players draw automatically
	fmt.Println()
	for i := 1; i < 4; i++ {
		fmt.Printf("Bot Player %d draws 5 cards...\n", i+1)
		for j := 0; j < cardsPerPlayer; j++ {
			h.deckService.DrawOne(i)
		}
	}

	// evaluate all hands
	players = h.gameService.GetPlayers()
	players = h.gameService.EvaluateAllHands(players)
	h.gameService.SavePlayers(players)

	// show results
	fmt.Println("\n=== RESULTS ===")
	for i, p := range players {
		label := fmt.Sprintf("%s (Bot)", p.Name)
		if i == 0 {
			label = fmt.Sprintf("%s (You)", p.Name)
		}
		fmt.Printf("%s: %s -> %s\n", label, formatCards(p.Cards), p.Hand.Name)
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

func (h *GameHandler) showPlayerMenu(playerIndex, cardNum, total int, scanner *bufio.Scanner) {
	for {
		fmt.Printf("[D] Draw card (%d/%d)  [S] Show current hand\n", cardNum, total)
		fmt.Print("> ")

		scanner.Scan()
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))

		switch input {
		case "d":
			card := h.deckService.DrawOne(playerIndex)
			fmt.Printf("  -> %s\n", card.String())
			fmt.Println()
			return
		case "s":
			players := h.gameService.GetPlayers()
			if len(players[playerIndex].Cards) == 0 {
				fmt.Println("  No cards yet.")
			} else {
				fmt.Printf("  Hand: %s\n", formatCards(players[playerIndex].Cards))
			}
			fmt.Println()
		default:
			fmt.Println("  Invalid input. Press D to draw or S to show hand.")
			fmt.Println()
		}
	}
}

func formatCards(cards []model.Card) string {
	var parts []string
	for _, c := range cards {
		parts = append(parts, c.String())
	}
	return strings.Join(parts, " ")
}
