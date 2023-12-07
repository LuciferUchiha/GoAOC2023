package main

import (
	"GoAOC2023/util"
	"fmt"
)

func main() {
	lines := util.ReadLines("day06/day06.in")
	timesAllowed := util.ExtractNumbers(lines[0])
	distancesToBeat := util.ExtractNumbers(lines[1])

	if len(timesAllowed) != len(distancesToBeat) {
		panic("Invalid input")
	}

	product := 1
	for raceIndex, timeAllowed := range timesAllowed {
		waysToWin := 0
		for i := 0; i <= timeAllowed; i++ {
			speed := i
			timeRemaining := timeAllowed - i
			distanceTravelled := speed * timeRemaining
			if distanceTravelled > distancesToBeat[raceIndex] {
				waysToWin++
			}
		}
		product *= waysToWin
	}
	fmt.Println(product)
}
