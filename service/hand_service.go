package service

import "sort"

type HandService struct{}

func NewHandService() *HandService {
	return &HandService{}
}

func countRanks(ranks []int) map[int]int {
	count := make(map[int]int)
	for _, r := range ranks {
		count[r]++
	}
	return count
}

func isFlush(suits []string) bool {
	for i := 1; i < len(suits); i++ {
		if suits[i] != suits[0] {
			return false
		}
	}
	return true
}

func isStraight(ranks []int) (bool, int) {
	sorted := make([]int, len(ranks))
	copy(sorted, ranks)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))

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

	// A-low straight: A-2-3-4-5 -> high card is 5, not A
	if sorted[0] == 14 && sorted[1] == 5 && sorted[2] == 4 && sorted[3] == 3 && sorted[4] == 2 {
		return true, 5
	}

	return false, 0
}

func (s *HandService) Evaluate(ranks []int, suits []string) (int, string, []int) {
	count := countRanks(ranks)
	flush := isFlush(suits)
	straight, straightHigh := isStraight(ranks)

	sorted := make([]int, len(ranks))
	copy(sorted, ranks)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))

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
	sort.Sort(sort.Reverse(sort.IntSlice(pairs)))
	sort.Sort(sort.Reverse(sort.IntSlice(singles)))

	if flush && straight && straightHigh == 14 {
		return 10, "Royal Flush", []int{14}
	}
	if flush && straight {
		return 9, "Straight Flush", []int{straightHigh}
	}
	if len(fours) == 1 {
		return 8, "Four of a Kind", append(fours, singles...)
	}
	if len(threes) == 1 && len(pairs) == 1 {
		return 7, "Full House", append(threes, pairs...)
	}
	if flush {
		return 6, "Flush", sorted
	}
	if straight {
		return 5, "Straight", []int{straightHigh}
	}
	if len(threes) == 1 {
		return 4, "Three of a Kind", append(threes, singles...)
	}
	if len(pairs) == 2 {
		return 3, "Two Pair", append(pairs, singles...)
	}
	if len(pairs) == 1 {
		return 2, "One Pair", append(pairs, singles...)
	}
	return 1, "High Card", sorted
}
