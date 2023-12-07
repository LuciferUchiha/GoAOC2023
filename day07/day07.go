package main

import (
	"GoAOC2023/util"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadLines("day07/day07.in")

	hands := make([][]rune, len(lines))
	bets := make([]int, len(lines))
	for i, line := range lines {
		tokens := strings.Split(line, " ")
		hand := tokens[0]
		bet := tokens[1]

		hands[i] = []rune(hand)
		betAmount, err := strconv.Atoi(bet)
		if err != nil {
			log.Fatal(err)
		}
		bets[i] = betAmount
	}

	for i, hand := range hands {
		fmt.Println(i, string(hand), bets[i], getHandRank(hand))
	}

	fmt.Println()

	// replace some of the cards so that they can easily be sorted via ascii value
	for i, hand := range hands {

		for j, card := range hand {
			hands[i][j] = replaceCard(card)
		}
	}

	originalHands := make([][]rune, len(hands))
	for i, hand := range hands {
		originalHands[i] = hand
	}

	// sort the hands by rank and then card values from left to right
	sort.Slice(hands, func(i, j int) bool {
		rankI := getHandRank(hands[i])
		rankJ := getHandRank(hands[j])
		if rankI == rankJ {
			for k := 0; k < len(hands[i]); k++ {
				if hands[i][k] == hands[j][k] {
					continue
				}
				return hands[i][k] < hands[j][k]
			}
		}
		return rankI < rankJ
	})

	// becuase the hands have been sorted, the bets need to be reordered to match the new hand order
	reordedBets := make([]int, len(bets))
	for i, hand := range hands {
		for j, originalHand := range originalHands {
			if string(hand) == string(originalHand) {
				reordedBets[i] = bets[j]
				break
			}
		}
	}

	for i, hand := range hands {
		fmt.Println(i, string(hand), reordedBets[i], getHandRank(hand))
	}

	sum := 0
	for i, _ := range hands {
		sum += reordedBets[i] * (i + 1)
	}
	fmt.Println(sum)
}

func replaceCard(card rune) rune {
	switch card {
	case 'T':
		return 'A'
	case 'J':
		return 'B'
	case 'Q':
		return 'C'
	case 'K':
		return 'D'
	case 'A':
		return 'E'
	default:
		return card
	}
}

func getHandRank(hand []rune) int {
	cardCounts := make(map[rune]int)
	for _, card := range hand {
		cardCounts[card]++
	}

	maxCount := 0
	for _, count := range cardCounts {
		if count > maxCount {
			maxCount = count
		}
	}
	// five of a kind
	if maxCount == 5 {
		return 6
	}
	// four of a kind
	if maxCount == 4 {
		return 5
	}

	if maxCount == 3 {
		// full house
		if len(cardCounts) == 2 {
			return 4
		}
		// three of a kind
		return 3
	}

	// two pair
	if len(cardCounts) == 3 {
		return 2
	}

	// one pair
	if len(cardCounts) == 4 {
		return 1
	}

	return 0
}
