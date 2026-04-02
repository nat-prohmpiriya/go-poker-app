package service

import "sort"

type HandService struct{}

func NewHandService() *HandService {
	return &HandService{}
}

// countRanks counts how many times each rank appears
// e.g. [K,K,K,10,10] -> {13:3, 10:2} = Full House
func countRanks(ranks []int) map[int]int {
	count := make(map[int]int)
	for _, r := range ranks {
		count[r]++
	}
	return count
}

// isFlush checks if all 5 cards have the same suit
func isFlush(suits []string) bool {
	for i := 1; i < len(suits); i++ {
		if suits[i] != suits[0] {
			return false
		}
	}
	return true
}

// isStraight checks if 5 cards are consecutive
// returns (isStraight, highCard)
// handles A-low straight: A-2-3-4-5 where high card is 5
func isStraight(ranks []int) (bool, int) {
	sorted := make([]int, len(ranks))
	copy(sorted, ranks)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))

	// normal case: e.g. 10-9-8-7-6
	straight := true
	for i := 1; i < len(sorted); i++ {
		if sorted[i-1]-sorted[i] != 1 {
			straight = false
			break
		}
	}
	if straight {
		return true, sorted[0]
	}

	// A-low case: A-2-3-4-5 -> sorted = [14,5,4,3,2]
	if sorted[0] == 14 && sorted[1] == 5 && sorted[2] == 4 && sorted[3] == 3 && sorted[4] == 2 {
		return true, 5 // high card is 5, not A
	}

	return false, 0
}

// Evaluate analyzes 5 cards and returns:
// - rank: hand strength 1-10 (10=Royal Flush, 1=High Card)
// - name: hand name e.g. "Flush", "One Pair"
// - tieBreak: values for comparing when hands have same rank
func (s *HandService) Evaluate(ranks []int, suits []string) (int, string, []int) {
	count := countRanks(ranks)
	flush := isFlush(suits)
	straight, straightHigh := isStraight(ranks)

	// sort ranks descending for tiebreak
	sorted := make([]int, len(ranks))
	copy(sorted, ranks)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))

	// group cards by count: fours(4), threes(3), pairs(2), singles(1)
	var fours, threes, pairs, singles []int
	for rank, cnt := range count {
		switch cnt {
		case 4:
			fours = append(fours, rank)
		case 3:
			threes = append(threes, rank)
		case 2:
			pairs = append(pairs, rank)
		default:
			singles = append(singles, rank)
		}
	}
	// sort descending for correct tiebreak order
	sort.Sort(sort.Reverse(sort.IntSlice(pairs)))
	sort.Sort(sort.Reverse(sort.IntSlice(singles)))

	// check from highest to lowest hand rank

	// 10: Royal Flush - flush + straight + high card is A
	if flush && straight && straightHigh == 14 {
		return 10, "Royal Flush", []int{14}
	}
	// 9: Straight Flush - flush + straight
	if flush && straight {
		return 9, "Straight Flush", []int{straightHigh}
	}
	// 8: Four of a Kind - 4 cards same rank
	if len(fours) == 1 {
		return 8, "Four of a Kind", append(fours, singles...)
	}
	// 7: Full House - 3 cards + 2 cards same rank
	if len(threes) == 1 && len(pairs) == 1 {
		return 7, "Full House", append(threes, pairs...)
	}
	// 6: Flush - all same suit
	if flush {
		return 6, "Flush", sorted
	}
	// 5: Straight - 5 consecutive ranks
	if straight {
		return 5, "Straight", []int{straightHigh}
	}
	// 4: Three of a Kind - 3 cards same rank
	if len(threes) == 1 {
		return 4, "Three of a Kind", append(threes, singles...)
	}
	// 3: Two Pair - 2 different pairs
	if len(pairs) == 2 {
		return 3, "Two Pair", append(pairs, singles...)
	}
	// 2: One Pair - 1 pair
	if len(pairs) == 1 {
		return 2, "One Pair", append(pairs, singles...)
	}
	// 1: High Card - nothing matched
	return 1, "High Card", sorted
}
