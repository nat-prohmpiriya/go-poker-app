package model

var Suits = []string{"Spades", "Hearts", "Diamonds", "Clubs"}

// Rank of card 2-14(J=11, Q=12,K=13, A=14)
var RankNames = map[int]string{
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "10",
	11: "J",
	12: "Q",
	13: "K",
	14: "A",
}

type Card struct {
	Rank int
	Suit string
}

func (c Card) String() string {
	return "[" + RankNames[c.Rank] + " " + c.Suit + "]"
}
