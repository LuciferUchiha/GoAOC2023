package main

import (
	"GoAOC2023/util"
	"fmt"
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
			if isGear(character) {
				surroundingNumbers := make([]int, 0)
				// copy it so that we can modify it without affecting the original for the bottom row case
				activeGrid := make([][]rune, len(grid))
				for i, line := range grid {
					activeGrid[i] = make([]rune, len(line))
					copy(activeGrid[i], line)
				}
				for _, lookupCoordinates := range surroundingLookup {
					lookupX := x + lookupCoordinates[0]
					lookupY := y + lookupCoordinates[1]
					// out of bounds
					if lookupX < 0 || lookupX >= len(line) || lookupY < 0 || lookupY >= len(activeGrid) {
						continue
					}
					lookupCharacter := activeGrid[lookupY][lookupX]
					if util.IsNumber(lookupCharacter) {
						// go to the left until no longer a number
						startOfNumberX := lookupX
						for startOfNumberX > 0 && util.IsNumber(activeGrid[lookupY][startOfNumberX-1]) {
							startOfNumberX--
						}
						// go to the right until no longer a number
						endOfNumberX := lookupX
						for endOfNumberX < len(line)-1 && util.IsNumber(activeGrid[lookupY][endOfNumberX+1]) {
							endOfNumberX++
						}

						// extract the number
						number := ""
						for i := startOfNumberX; i <= endOfNumberX; i++ {
							number += string(activeGrid[lookupY][i])
						}
						convertedNumber, err := strconv.Atoi(number)
						if err != nil {
							panic(err)
						}

						// replace the number with dots
						for i := startOfNumberX; i <= endOfNumberX; i++ {
							activeGrid[lookupY][i] = '.'
						}

						fmt.Println("surrounding number:", convertedNumber)
						surroundingNumbers = append(surroundingNumbers, convertedNumber)
					}
				}
				if len(surroundingNumbers) == 2 {
					sum += surroundingNumbers[0] * surroundingNumbers[1]
				}
			}
		}
	}
	fmt.Println(sum)
}

func isGear(symbol rune) bool {
	// matches all special characters except for .
	return symbol == '*'
}
