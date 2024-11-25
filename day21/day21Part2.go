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

	fieldHeight := len(field)
	fieldWidth := len(field[0])

	startPos := util.Point{X: 0, Y: 0}

	for y, row := range field {
		for x, symbol := range row {
			if symbol == Start {
				startPos.X = x
				startPos.Y = y
			}
		}
	}

	prevMap := make([][]int, fieldHeight)
	currentMap := make([][]int, fieldHeight)
	for i := range prevMap {
		prevMap[i] = make([]int, fieldWidth)
		currentMap[i] = make([]int, fieldWidth)
	}

	prevMap[startPos.Y][startPos.X] = 1

	// replace start position with empty space
	field[startPos.Y][startPos.X] = Empty

	fmt.Println(startPos)

	steps := 5000

	adjacent := []util.Point{
		{X: 0, Y: 1},
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: -1, Y: 0},
	}

	for i := 0; i < steps; i++ {
		fmt.Println("Step", i)
		for y, row := range prevMap {
			for x, val := range row {
				if val == 0 {
					continue
				}
				for _, adj := range adjacent {
					newPos := util.Point{X: x + adj.X, Y: y + adj.Y}

					mapPos := util.Point{X: newPos.X, Y: newPos.Y}
					for mapPos.X < 0 {
						mapPos.X += fieldWidth
					}
					for mapPos.Y < 0 {
						mapPos.Y += fieldHeight
					}
					mapPos.X %= fieldWidth
					mapPos.Y %= fieldHeight

					if field[mapPos.Y][mapPos.X] == Wall {
						continue
					}
					//fmt.Println("Setting", mapPos, "for", newPos, "to", val)
					if currentMap[mapPos.Y][mapPos.X] == 0 {
						currentMap[mapPos.Y][mapPos.X] += val
					}
				}
			}
		}
		// copy currentMap to prevMap, reset currentMap
		for y, row := range currentMap {
			for x, val := range row {
				prevMap[y][x] = val
				currentMap[y][x] = 0
			}
		}
	}

	sum := 0
	for _, row := range prevMap {
		for _, val := range row {
			sum += val
		}
	}
	println(sum)
}
