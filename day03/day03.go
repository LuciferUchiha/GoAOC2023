package main

import (
	"GoAOC2023/util"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines := util.ReadLines("day03/day03.in")

	sum := 0
	// transform lines into a 2D array of runes
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	surroundingLookup := [][]int{
		{-1, -1}, // top left
		{0, -1},  // top
		{1, -1},  // top right
		{-1, 0},  // left
		{1, 0},   // right
		{-1, 1},  // bottom left
		{0, 1},   // bottom
		{1, 1},   // bottom right
	}

	for y, line := range grid {
		for x, character := range line {
			if isSymbol(character) {
				// look at the 8 surrounding characters are they a number or part of a number?
				// if so, add the full number to the sum and replace the number with dots
				for _, lookupCoordinates := range surroundingLookup {
					lookupX := x + lookupCoordinates[0]
					lookupY := y + lookupCoordinates[1]
					// out of bounds
					if lookupX < 0 || lookupX >= len(line) || lookupY < 0 || lookupY >= len(grid) {
						continue
					}
					lookupCharacter := grid[lookupY][lookupX]
					if util.IsNumber(lookupCharacter) {
						// go to the left until no longer a number
						startOfNumberX := lookupX
						for startOfNumberX > 0 && util.IsNumber(grid[lookupY][startOfNumberX-1]) {
							startOfNumberX--
						}
						// go to the right until no longer a number
						endOfNumberX := lookupX
						for endOfNumberX < len(line)-1 && util.IsNumber(grid[lookupY][endOfNumberX+1]) {
							endOfNumberX++
						}

						// extract the number
						number := ""
						for i := startOfNumberX; i <= endOfNumberX; i++ {
							number += string(grid[lookupY][i])
						}
						convertedNumber, err := strconv.Atoi(number)
						if err != nil {
							panic(err)
						}
						sum += convertedNumber

						// replace the number with dots
						for i := startOfNumberX; i <= endOfNumberX; i++ {
							grid[lookupY][i] = '.'
						}
					}
				}
			}
		}
	}
	fmt.Println(sum)
}

func isSymbol(symbol rune) bool {
	// matches all special characters except for .
	re := regexp.MustCompile("[^a-zA-Z0-9\\s.]")
	return re.MatchString(string(symbol))
}
