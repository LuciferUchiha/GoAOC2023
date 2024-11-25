package main

import (
	"GoAOC2023/util"
	"fmt"
	"math"
)

func main() {
	lines := util.ReadLines("day11/day11.in")
	space := make([][]rune, len(lines))
	for i, line := range lines {
		space[i] = []rune(line)
		fmt.Println(space[i])
	}
	spaceWidth := len(space[0])
	spaceHeight := len(space)

	galaxy := '#'

	galaxies := make([]util.Point, 0)
	for y, row := range space {
		for x, char := range row {
			if char == galaxy {
				galaxies = append(galaxies, util.Point{X: x, Y: y})
			}
		}
	}
	emptyRows := make([]int, 0)
	emptyCols := make([]int, 0)
	for y, row := range space {
		empty := true
		for _, char := range row {
			if char == galaxy {
				empty = false
				break
			}
		}
		if empty {
			emptyRows = append(emptyRows, y)
		}
	}
	for x := 0; x < spaceWidth; x++ {
		empty := true
		for y := 0; y < spaceHeight; y++ {
			if space[y][x] == galaxy {
				empty = false
				break
			}
		}
		if empty {
			emptyCols = append(emptyCols, x)
		}
	}
	fmt.Println("Empty rows:", emptyRows)
	fmt.Println("Empty cols:", emptyCols)

	expandedGalaxies := make([]util.Point, 0)
	for _, galaxy := range galaxies {
		expandedGalaxies = append(expandedGalaxies, util.Point{X: galaxy.X, Y: galaxy.Y})
	}

	fmt.Println("Original galaxies:", galaxies)
	rateOfExpansion := 999999 // for part 1 this is 1 for part 2 this is 999999
	for _, emptyRow := range emptyRows {
		// increase all y values of galaxies after emptyRow by 1
		for i := range galaxies {
			if galaxies[i].Y > emptyRow {
				expandedGalaxies[i].Y += rateOfExpansion
			}
		}
	}

	for _, emptyCol := range emptyCols {
		// increase all x values of galaxies after emptyCol by 1
		for i := range galaxies {
			if galaxies[i].X > emptyCol {
				expandedGalaxies[i].X += rateOfExpansion
			}
		}
	}

	fmt.Println("Updated galaxies:", expandedGalaxies)

	totalDistance := 0
	for i := 0; i < len(expandedGalaxies); i++ {
		for j := i + 1; j < len(expandedGalaxies); j++ {
			distance := int(math.Floor(util.ManhattanDistance(expandedGalaxies[i], expandedGalaxies[j])))
			fmt.Printf("Distance between %v and %v is %d\n", expandedGalaxies[i], expandedGalaxies[j], distance)
			totalDistance += distance
		}
	}
	fmt.Println("Total distance:", totalDistance)
}
