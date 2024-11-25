package main

import (
	"GoAOC2023/util"
	"fmt"
)

const rock = 'O'
const wall = '#'
const empty = '.'

func main() {
	lines := util.ReadLines("day14/day14.in")

	field := make([][]rune, 0)
	for _, line := range lines {
		field = append(field, []rune(line))
	}
	fmt.Println("Original field")
	printField(field)

	// do 1000000000 cycles, we expect the field to repeat itself after a certain amount of cycles hence only 1000
	for i := 0; i < 1_000; i++ {
		println("cycle", i)
		// copy field
		field = tiltNorth(field)
		field = tiltWest(field)
		field = tiltSouth(field)
		field = tiltEast(field)
	}

	fmt.Println("Cycled field")
	printField(field)

	load := getLoadOnNorthBeam(field)
	fmt.Println("Load on north beam", load)
}

func tiltNorth(field [][]rune) [][]rune {
	fieldHeight := len(field)
	fieldWidth := len(field[0])
	// we start at the top left and slowly move them up
	for y := 0; y < fieldHeight; y++ {
		for x := 0; x < fieldWidth; x++ {
			if field[y][x] == rock {
				field[y][x] = empty
				var currentY = y
				// as long as above there is space it moves north
				for currentY > 0 && field[currentY-1][x] == empty {
					currentY--
				}
				field[currentY][x] = rock
			}
		}
	}
	return field
}

func tiltEast(field [][]rune) [][]rune {
	// we start at the bottom right and slowly move them right
	fieldHeight := len(field)
	fieldWidth := len(field[0])
	for y := fieldHeight - 1; y >= 0; y-- {
		for x := fieldWidth - 1; x >= 0; x-- {
			if field[y][x] == rock {
				field[y][x] = empty
				var currentX = x
				// as long as right there is space it moves east
				for currentX < fieldWidth-1 && field[y][currentX+1] == empty {
					currentX++
				}
				field[y][currentX] = rock
			}
		}
	}
	return field
}

func tiltSouth(field [][]rune) [][]rune {
	// we start at the bottom right and slowly move them down
	fieldHeight := len(field)
	fieldWidth := len(field[0])
	for y := fieldHeight - 1; y >= 0; y-- {
		for x := fieldWidth - 1; x >= 0; x-- {
			if field[y][x] == rock {
				field[y][x] = empty
				var currentY = y
				// as long as below there is space it moves south
				for currentY < fieldHeight-1 && field[currentY+1][x] == empty {
					currentY++
				}
				field[currentY][x] = rock
			}
		}
	}
	return field
}

func tiltWest(field [][]rune) [][]rune {
	// we start at the top left and slowly move them left
	fieldHeight := len(field)
	fieldWidth := len(field[0])
	for y := 0; y < fieldHeight; y++ {
		for x := 0; x < fieldWidth; x++ {
			if field[y][x] == rock {
				field[y][x] = empty
				var currentX = x
				// as long as left there is space it moves west
				for currentX > 0 && field[y][currentX-1] == empty {
					currentX--
				}
				field[y][currentX] = rock
			}
		}
	}
	return field
}

func getLoadOnNorthBeam(field [][]rune) int {
	fieldHeight := len(field)
	load := 0
	for y, row := range field {
		for _, symbol := range row {
			if symbol == rock {
				load += fieldHeight - y
				fmt.Println("added", fieldHeight-y)
			}
		}
	}
	return load
}

func printField(field [][]rune) {
	for _, row := range field {
		println(string(row))
	}
}
