package service

import (
	"testing"
)

func TestEvaluateRoyalFlush(t *testing.T) {
	s := NewHandService()
	ranks := []int{14, 13, 12, 11, 10}
	suits := []string{"Spades", "Spades", "Spades", "Spades", "Spades"}

	rank, name, _ := s.Evaluate(ranks, suits)

	if rank != 10 || name != "Royal Flush" {
		t.Errorf("expected Royal Flush (10), got %s (%d)", name, rank)
	}
}

func TestEvaluateStraightFlush(t *testing.T) {
	s := NewHandService()
	ranks := []int{9, 8, 7, 6, 5}
	suits := []string{"Hearts", "Hearts", "Hearts", "Hearts", "Hearts"}

	rank, name, tb := s.Evaluate(ranks, suits)

	if rank != 9 || name != "Straight Flush" {
		t.Errorf("expected Straight Flush (9), got %s (%d)", name, rank)
	}
	if tb[0] != 9 {
		t.Errorf("expected high card 9, got %d", tb[0])
	}
}

func TestEvaluateFourOfAKind(t *testing.T) {
	s := NewHandService()
	ranks := []int{8, 8, 8, 8, 12}
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs", "Spades"}

	rank, name, tb := s.Evaluate(ranks, suits)

	if rank != 8 || name != "Four of a Kind" {
		t.Errorf("expected Four of a Kind (8), got %s (%d)", name, rank)
	}
	if tb[0] != 8 || tb[1] != 12 {
		t.Errorf("expected tiebreak [8, 12], got %v", tb)
	}
}

func TestEvaluateFullHouse(t *testing.T) {
	s := NewHandService()
	ranks := []int{13, 13, 13, 10, 10}
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades", "Hearts"}

	rank, name, tb := s.Evaluate(ranks, suits)

	if rank != 7 || name != "Full House" {
		t.Errorf("expected Full House (7), got %s (%d)", name, rank)
	}
	if tb[0] != 13 || tb[1] != 10 {
		t.Errorf("expected tiebreak [13, 10], got %v", tb)
	}
}

func TestEvaluateFlush(t *testing.T) {
	s := NewHandService()
	ranks := []int{14, 10, 5, 3, 2}
	suits := []string{"Spades", "Spades", "Spades", "Spades", "Spades"}

	rank, name, tb := s.Evaluate(ranks, suits)

	if rank != 6 || name != "Flush" {
		t.Errorf("expected Flush (6), got %s (%d)", name, rank)
	}
	if tb[0] != 14 {
		t.Errorf("expected high card 14, got %d", tb[0])
	}
}

func TestEvaluateStraight(t *testing.T) {
	s := NewHandService()
	ranks := []int{10, 9, 8, 7, 6}
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs", "Spades"}

	rank, name, tb := s.Evaluate(ranks, suits)

	if rank != 5 || name != "Straight" {
		t.Errorf("expected Straight (5), got %s (%d)", name, rank)
	}
	if tb[0] != 10 {
		t.Errorf("expected high card 10, got %d", tb[0])
	}
}

func TestEvaluateStraightALow(t *testing.T) {
	s := NewHandService()
	ranks := []int{14, 2, 3, 4, 5}
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs", "Spades"}

	rank, name, tb := s.Evaluate(ranks, suits)

	if rank != 5 || name != "Straight" {
		t.Errorf("expected Straight (5), got %s (%d)", name, rank)
	}
	if tb[0] != 5 {
		t.Errorf("expected high card 5 (A-low), got %d", tb[0])
	}
}

func TestEvaluateThreeOfAKind(t *testing.T) {
	s := NewHandService()
	ranks := []int{12, 12, 12, 7, 3}
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs", "Spades"}

	rank, name, tb := s.Evaluate(ranks, suits)

	if rank != 4 || name != "Three of a Kind" {
		t.Errorf("expected Three of a Kind (4), got %s (%d)", name, rank)
	}
	if tb[0] != 12 {
		t.Errorf("expected three rank 12, got %d", tb[0])
	}
}

func TestEvaluateTwoPair(t *testing.T) {
	s := NewHandService()
	ranks := []int{13, 13, 9, 9, 14}
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs", "Spades"}

	rank, name, tb := s.Evaluate(ranks, suits)

	if rank != 3 || name != "Two Pair" {
		t.Errorf("expected Two Pair (3), got %s (%d)", name, rank)
	}
	if tb[0] != 13 || tb[1] != 9 || tb[2] != 14 {
		t.Errorf("expected tiebreak [13, 9, 14], got %v", tb)
	}
}

func TestEvaluateOnePair(t *testing.T) {
	s := NewHandService()
	ranks := []int{11, 11, 14, 8, 4}
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs", "Spades"}

	rank, name, tb := s.Evaluate(ranks, suits)

	if rank != 2 || name != "One Pair" {
		t.Errorf("expected One Pair (2), got %s (%d)", name, rank)
	}
	if tb[0] != 11 || tb[1] != 14 || tb[2] != 8 || tb[3] != 4 {
		t.Errorf("expected tiebreak [11, 14, 8, 4], got %v", tb)
	}
}

func TestEvaluateHighCard(t *testing.T) {
	s := NewHandService()
	ranks := []int{14, 13, 9, 5, 2}
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs", "Spades"}

	rank, name, tb := s.Evaluate(ranks, suits)

	if rank != 1 || name != "High Card" {
		t.Errorf("expected High Card (1), got %s (%d)", name, rank)
	}
	if tb[0] != 14 || tb[1] != 13 || tb[2] != 9 || tb[3] != 5 || tb[4] != 2 {
		t.Errorf("expected tiebreak [14, 13, 9, 5, 2], got %v", tb)
	}
}
