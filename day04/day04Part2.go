package main

import (
	"GoAOC2023/util"
	"fmt"
	"strings"
)

func main() {
	lines := util.ReadLines("day04/day04.in")
	// add all the copies
	originalCards := make([]string, len(lines))

	for i, line := range lines {
		originalCards[i] = line
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		cardId := getCardId(line)
		winningNumbers := getWinningNumbers(line)
		myNumbers := getMyNumbers(line)

		cardScore := 0
		for _, number := range myNumbers {
			if util.ArrayContains(winningNumbers, number) {
				cardScore++
			}
		}

		for i := 0; i < cardScore; i++ {
			// add copy of line to lines
			copyCard := originalCards[cardId+i]
			lines = append(lines, copyCard)
		}
	}

	fmt.Println(len(lines))
}

func getWinningNumbers(line string) []int {
	numbers := strings.Split(line, ":")[1]
	sections := strings.Split(numbers, "|")
	sections[0] = strings.TrimSpace(sections[0])
	winningNumbersTokens := strings.Split(sections[0], " ")
	winningNumbersTokens = util.StringArrayTrimElements(winningNumbersTokens)
	winningNumbersTokens = util.StringArrayRemoveEmptyStrings(winningNumbersTokens)
	winningNumbers := util.StringArrayToIntArray(winningNumbersTokens)
	return winningNumbers
}

func getMyNumbers(line string) []int {
	numbers := strings.Split(line, ":")[1]
	sections := strings.Split(numbers, "|")
	sections[1] = strings.TrimSpace(sections[1])
	myNumbersTokens := strings.Split(sections[1], " ")
	myNumbersTokens = util.StringArrayTrimElements(myNumbersTokens)
	myNumbersTokens = util.StringArrayRemoveEmptyStrings(myNumbersTokens)
	myNumbers := util.StringArrayToIntArray(myNumbersTokens)
	return myNumbers
}

func getCardId(line string) int {
	cardId := util.ExtractNumber(strings.Split(line, ":")[0])
	return cardId
}
