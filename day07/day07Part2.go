package main

import (
	"GoAOC2023/util"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

const NOTHING = 0
const PAIR = 1
const TWO_PAIR = 2
const THREE_OF_A_KIND = 3
const FULL_HOUSE = 4
const FOUR_OF_A_KIND = 5
const FIVE_OF_A_KIND = 6

const JOKER = '1'

type Hand struct {
	cards []rune
	bet   int
}

func main() {
	lines := util.ReadLines("day07/day07.in")

	hands := make([]Hand, len(lines))
	for i, line := range lines {
		tokens := strings.Split(line, " ")
		hand := tokens[0]
		bet := tokens[1]

		handArray := []rune(hand)
		betAmount, err := strconv.Atoi(bet)
		if err != nil {
			log.Fatal(err)
		}
		hands[i] = Hand{handArray, betAmount}
	}

	// replace some of the cards so that they can easily be sorted via ascii value
	for i, hand := range hands {

		for j, card := range hand.cards {
			hands[i].cards[j] = replaceCard(card)
		}
	}

	// sort the hands by rank and then card values from left to right
	sort.Slice(hands, func(i, j int) bool {
		rankI := getHandRank(hands[i].cards)
		rankJ := getHandRank(hands[j].cards)
		if rankI == rankJ {
			for k := 0; k < len(hands[i].cards); k++ {
				if hands[i].cards[k] == hands[j].cards[k] {
					continue
				}
				return hands[i].cards[k] < hands[j].cards[k]
			}
		}
		return rankI < rankJ
	})

	for i, hand := range hands {
		handString := ""
		for _, card := range hands[i].cards {
			handString += string(convertBack(card))
		}
		fmt.Println(i, handString, hand.bet, getHandRank(hand.cards))
	}

	sum := 0
	for i, hand := range hands {
		sum += hand.bet * (i + 1)
	}
	fmt.Println(sum)
}

func replaceCard(card rune) rune {
	switch card {
	case 'J':
		return JOKER // our new joker that is worth the least
	case 'T':
		return 'A'
	case 'Q':
		return 'B'
	case 'K':
		return 'C'
	case 'A':
		return 'D'
	default:
		return card
	}
}

func convertBack(card rune) rune {
	switch card {
	case JOKER:
		return 'J'
	case 'A':
		return 'T'
	case 'B':
		return 'Q'
	case 'C':
		return 'K'
	case 'D':
		return 'A'
	default:
		return card
	}
}

func getHandRank(hand []rune) int {
	cardCounts := make(map[rune]int)
	for _, card := range hand {
		cardCounts[card]++
	}

	numberOfJokers := cardCounts[JOKER]

	maxCount := 0
	for _, count := range cardCounts {
		if count > maxCount {
			maxCount = count
		}
	}

	if maxCount == 5 {
		return FIVE_OF_A_KIND
	}

	if maxCount == 4 {
		if numberOfJokers > 0 {
			return FIVE_OF_A_KIND
		}
		return FOUR_OF_A_KIND
	}

	if maxCount == 3 {
		if len(cardCounts) == 2 {
			// 3 and 2
			if numberOfJokers > 0 {
				// can be turned into a five of a kind
				return FIVE_OF_A_KIND
			}
			return FULL_HOUSE
		}
		// three and 2 different cards
		if numberOfJokers > 0 {
			// can be a four of a kind
			return FOUR_OF_A_KIND
		}
		return THREE_OF_A_KIND
	}

	if len(cardCounts) == 3 {
		if numberOfJokers == 2 {
			// is one of the pairs can make four of a kind
			return FOUR_OF_A_KIND
		}
		if numberOfJokers == 1 {
			// is the kicker, can make a full house
			return FULL_HOUSE
		}
		return TWO_PAIR
	}

	if len(cardCounts) == 4 {
		if numberOfJokers > 0 {
			// can be a three of a kind, is better than making two pairs
			return THREE_OF_A_KIND
		}
		return PAIR
	}

	if numberOfJokers > 0 {
		return PAIR
	}

	return NOTHING
}
