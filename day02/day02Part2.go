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

	for _, line := range lines {
		sum += powerOfMinimumGame(line)
	}
	fmt.Println(sum)
}

func powerOfMinimumGame(line string) int {
	re := regexp.MustCompile(`(?P<Count>\d+) (?P<Color>\w+)`)

	matches := re.FindAllStringSubmatch(line, -1)

	maxRedCubes := 0
	maxGreenCubes := 0
	maxBlueCubes := 0
	for _, match := range matches {
		countStr := match[re.SubexpIndex("Count")]
		color := match[re.SubexpIndex("Color")]

		count, err := strconv.Atoi(countStr)
		if err != nil {
			fmt.Println("Error converting count to int:", err)
			return -1
		}

		switch color {
		case "red":
			if count > maxRedCubes {
				maxRedCubes = count
			}
		case "green":
			if count > maxGreenCubes {
				maxGreenCubes = count
			}
		case "blue":
			if count > maxBlueCubes {
				maxBlueCubes = count
			}
		default:
			fmt.Println("Unknown color:", color)
			return -1
		}
	}

	return maxRedCubes * maxGreenCubes * maxBlueCubes
}
