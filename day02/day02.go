package main

import (
	"GoAOC2023/util"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines := util.ReadLines("day02/day02.in")
	sum := 0
	numRedCubes := 12
	numGreenCubes := 13
	numBlueCubes := 14

	for gameId, line := range lines {
		if isGamePossible(line, numRedCubes, numGreenCubes, numBlueCubes) {
			sum += gameId + 1
		}
	}
	fmt.Println(sum)
}

func isGamePossible(line string, numRedCubes int, numGreenCubes int, numBlueCubes int) bool {
	re := regexp.MustCompile(`(?P<Count>\d+) (?P<Color>\w+)`)

	matches := re.FindAllStringSubmatch(line, -1)
	for _, match := range matches {
		countStr := match[re.SubexpIndex("Count")]
		color := match[re.SubexpIndex("Color")]

		count, err := strconv.Atoi(countStr)
		if err != nil {
			fmt.Println("Error converting count to int:", err)
			return false
		}

		switch color {
		case "red":
			if count > numRedCubes {
				return false
			}
		case "green":
			if count > numGreenCubes {
				return false
			}
		case "blue":
			if count > numBlueCubes {
				return false
			}
		default:
			fmt.Println("Unknown color:", color)
			return false
		}
	}

	return true
}
