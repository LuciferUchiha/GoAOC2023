package main

import (
	"GoAOC2023/util"
	"fmt"
)

func main() {
	lines := util.ReadLines("day06/day06.in")
	timesAllowed := util.ExtractNumbers(lines[0])
	distancesToBeat := util.ExtractNumbers(lines[1])

	actualTimeAllowed := util.ConcatNumbersInIntArr(timesAllowed)
	actualDistancesToBeat := util.ConcatNumbersInIntArr(distancesToBeat)

	waysToWin := 0
	for i := 0; i <= actualTimeAllowed; i++ {
		speed := i
		timeRemaining := actualTimeAllowed - i
		distanceTravelled := speed * timeRemaining
		if distanceTravelled > actualDistancesToBeat {
			waysToWin++
		}
	}
	fmt.Println(waysToWin)
}
