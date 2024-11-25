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

	tiltedField := tiltNorth(field)
	fmt.Println("Tilted field")
	printField(tiltedField)

	load := getLoadOnNorthBeam(tiltedField)
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
