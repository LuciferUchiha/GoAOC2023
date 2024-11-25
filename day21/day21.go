package main

import (
	"GoAOC2023/util"
	"fmt"
)

const (
	Wall  = '#'
	Empty = '.'
	Start = 'S'
)

func main() {
	lines := util.ReadLines("day21/day21.in")

	field := make([][]rune, len(lines))
	for i, line := range lines {
		field[i] = []rune(line)
	}

	startPos := util.Point{X: 0, Y: 0}
	for y, row := range field {
		for x, c := range row {
			if c == Start {
				startPos = util.Point{X: x, Y: y}
			}
		}
	}

	fmt.Println(startPos)

	steps := 64
	lastPositions := map[util.Point]bool{startPos: true}
	currentPositions := map[util.Point]bool{}

	adjacent := []util.Point{
		{X: 0, Y: 1},
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: -1, Y: 0},
	}

	for i := 0; i < steps; i++ {
		for pos := range lastPositions {
			for _, adj := range adjacent {
				newPos := pos.Add(adj)

				if field[newPos.Y][newPos.X] == Wall {
					continue
				}

				if _, ok := currentPositions[newPos]; ok {
					continue
				}
				currentPositions[newPos] = true
			}
		}
		lastPositions = currentPositions
		currentPositions = map[util.Point]bool{}
	}

	println(len(lastPositions))
	for pos := range lastPositions {
		fmt.Println(pos)
	}
}
