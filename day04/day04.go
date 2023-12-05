package main

import (
	"GoAOC2023/util"
	"fmt"
	"math"
	"strings"
)

func main() {
	lines := util.ReadLines("day04/day04.in")
	sum := 0
	for _, line := range lines {
		numbers := strings.Split(line, ":")[1]
		sections := strings.Split(numbers, "|")
		sections[0] = strings.TrimSpace(sections[0])
		sections[1] = strings.TrimSpace(sections[1])
		winningNumbersTokens := strings.Split(sections[0], " ")
		winningNumbersTokens = util.StringArrayTrimElements(winningNumbersTokens)
		winningNumbersTokens = util.StringArrayRemoveEmptyStrings(winningNumbersTokens)
		winningNumbers := util.StringArrayToIntArray(winningNumbersTokens)
		myNumbersTokens := strings.Split(sections[1], " ")
		myNumbersTokens = util.StringArrayTrimElements(myNumbersTokens)
		myNumbersTokens = util.StringArrayRemoveEmptyStrings(myNumbersTokens)
		myNumbers := util.StringArrayToIntArray(myNumbersTokens)

		cardScore := 0
		for _, number := range myNumbers {
			if util.ArrayContains(winningNumbers, number) {
				cardScore++
			}
		}
		points := int(math.Pow(2, float64(cardScore-1)))
		sum += points
	}
	fmt.Println(sum)
}
