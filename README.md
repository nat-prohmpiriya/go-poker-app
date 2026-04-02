# Poker Hand Evaluator

Poker card dealing and hand evaluation system built with Go. Simulates a 4-player poker game with deck creation, shuffling, dealing, hand evaluation, and winner determination with full tie-breaking support.

## Branches

| Branch | Mode | Description |
|--------|------|-------------|
| `main` | Auto | Run once, deal cards to 4 players, evaluate hands, print winner |
| `interactive-cli` | Interactive | Player 1 draws cards manually (press D), Player 2-4 are bots |

## How to Run

```bash
cd poker-app
go run .
```

### main branch — Auto mode

```
Player 1: [A Spades] [10 Spades] [5 Spades] [2 Spades] [K Spades] -> Flush
Player 2: [J Diamonds] [J Clubs] [4 Hearts] [9 Spades] [7 Clubs] -> One Pair
Player 3: [Q Hearts] [8 Spades] [2 Clubs] [6 Diamonds] [A Clubs] -> High Card
Player 4: [K Hearts] [K Diamonds] [K Clubs] [10 Spades] [10 Hearts] -> Full House

Cards left in deck: 32

*** Winner is Player 4 with Full House! ***
```

### interactive-cli branch — Interactive mode

```
=== POKER GAME ===
[1] Start New Round
[2] Exit

> 1

--- Your turn (Player 1) ---
[D] Draw card (1/5)  [S] Show current hand
> d
  -> [A Spades]

[D] Draw card (2/5)  [S] Show current hand
> d
  -> [10 Hearts]
...

Bot Player 2 draws 5 cards...
Bot Player 3 draws 5 cards...
Bot Player 4 draws 5 cards...

=== RESULTS ===
Player 1 (You): [A Spades] [10 Hearts] ... -> High Card
Player 2 (Bot): [J Diamonds] [J Clubs] ... -> One Pair
...

*** Winner is Player 2 with One Pair! ***
```

## Project Structure

```
poker-app/
├── main.go
├── handler/
│   └── game_handler.go     # game flow control + output
├── service/
│   ├── deck_service.go     # create deck, shuffle, deal, draw
│   ├── hand_service.go     # evaluate poker hands (core logic)
│   └── game_service.go     # evaluate all players + determine winner
├── repository/
│   └── game_repository.go  # in-memory state storage
├── model/
│   ├── card.go             # Card struct (Rank, Suit)
│   ├── deck.go             # Deck struct
│   ├── hand.go             # Hand evaluation result
│   └── player.go           # Player struct
```

## Architecture

```
main → handler → service → repository → model
```

| Layer | Responsibility |
|-------|---------------|
| Handler | Controls game flow, formats output |
| Service | Business logic — shuffle, evaluate, compare |
| Repository | In-memory state management |
| Model | Data structures |

## Hand Rankings (high to low)

| Rank | Hand | Example |
|------|------|---------|
| 10 | Royal Flush | A-K-Q-J-10 same suit |
| 9 | Straight Flush | 9-8-7-6-5 same suit |
| 8 | Four of a Kind | K-K-K-K-3 |
| 7 | Full House | K-K-K-10-10 |
| 6 | Flush | A-10-5-3-2 same suit |
| 5 | Straight | 10-9-8-7-6 |
| 4 | Three of a Kind | Q-Q-Q-7-3 |
| 3 | Two Pair | K-K-9-9-A |
| 2 | One Pair | J-J-A-8-4 |
| 1 | High Card | A-K-9-5-2 |

## Tech Stack

- Go (Golang)
- No external dependencies
- No API / No Database — CLI only, in-memory
