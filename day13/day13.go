package main

import (
	"GoAOC2023/util"
	"fmt"
	"strings"
)

func main() {
	lines := util.ReadLines("day13/day13.in")

	patterns := make([][]string, 0)
	currentPattern := make([]string, 0)
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			patterns = append(patterns, currentPattern)
			currentPattern = make([]string, 0)
		} else {
			currentPattern = append(currentPattern, line)
		}
	}
	patterns = append(patterns, currentPattern)

	score := 0
	for i, pattern := range patterns {
		fmt.Println("Original pattern", i)
		for _, row := range pattern {
			fmt.Println(row)
		}
		originalMirror := util.Point{}
		originalMirrorIsHorizontal := false

		originalIsHorizontalMirror, originalStartingPoint := getHorizontalMirror(pattern)
		if originalIsHorizontalMirror {
			originalMirror = originalStartingPoint
			originalMirrorIsHorizontal = true
		}

		// check if it is a vertical mirror
		originalIsVerticalMirror, originalStartingPoint := getVerticalMirror(pattern)
		if originalIsVerticalMirror {
			originalMirror = originalStartingPoint
		}

		fmt.Println("Original mirror", originalMirror, "horizontal =", originalMirrorIsHorizontal)

		for y, row := range pattern {
			for x, symbol := range row {
				// flip the symbol
				patternCopy := make([]string, len(pattern))
				copy(patternCopy, pattern)
				if symbol == '.' {
					patternCopy[y] = patternCopy[y][:x] + "#" + patternCopy[y][x+1:]
				} else {
					patternCopy[y] = patternCopy[y][:x] + "." + patternCopy[y][x+1:]
				}

				fmt.Println("Flipped at ", x, y)

				// check if it is a horizontal mirror
				isHorizontalMirror, startingPoint := getHorizontalMirror(patternCopy)
				fmt.Println("Horizontal mirror", isHorizontalMirror, "starting point", startingPoint)
				if isHorizontalMirror && (startingPoint != originalMirror || !originalMirrorIsHorizontal) {
					fmt.Println("Different Horizontal mirror")
					fmt.Println("At rows", startingPoint)
					fmt.Println("Flipped", x, y)
					score += 100 * (startingPoint.X + 1)
					goto nextPattern
				}

				// check if it is a vertical mirror
				isVerticalMirror, startingPoint := getVerticalMirror(patternCopy)
				fmt.Println("Vertical mirror", isVerticalMirror, "starting point", startingPoint)
				if isVerticalMirror && (startingPoint != originalMirror || originalMirrorIsHorizontal) {
					fmt.Println("Different Vertical mirror")
					fmt.Println("At columns", startingPoint)
					fmt.Println("Flipped", x, y)
					score += startingPoint.X + 1
					goto nextPattern
				}
			}
		}
		fmt.Println("No mirror found")
	nextPattern:
	}
	fmt.Println("Score:", score)
}

func getHorizontalMirror(pattern []string) (bool, util.Point) {
	patternHeight := len(pattern)

	matchingRows := make([]util.Point, 0)
	for i, row := range pattern {
		for j := i + 1; j < patternHeight; j++ {
			otherRow := pattern[j]
			if row == otherRow {
				matchingRows = append(matchingRows, util.Point{X: i, Y: j})
			}
		}
	}

	potentialStartingPoints := make([]util.Point, 0)
	for _, matchingRow := range matchingRows {
		// Y is always greater than X
		if matchingRow.X+1 == matchingRow.Y {
			potentialStartingPoints = append(potentialStartingPoints, matchingRow)
		}
	}
	for _, startingPoint := range potentialStartingPoints {
		leftPointer := startingPoint.X
		rightPointer := startingPoint.Y
		// until one of the pointers is out of bounds
		for leftPointer >= 0 && rightPointer < patternHeight {
			if pattern[leftPointer] != pattern[rightPointer] {
				break
			}
			leftPointer--
			rightPointer++
		}
		if leftPointer < 0 || rightPointer >= patternHeight {
			return true, startingPoint
		}
	}
	return false, util.Point{}
}

func getVerticalMirror(pattern []string) (bool, util.Point) {
	patternWidth := len(pattern[0])

	matchingColumns := make([]util.Point, 0)
	for i := 0; i < patternWidth; i++ {
		for j := i + 1; j < patternWidth; j++ {
			column := getColumn(pattern, i)
			otherColumn := getColumn(pattern, j)
			if column == otherColumn {
				matchingColumns = append(matchingColumns, util.Point{X: i, Y: j})
			}
		}
	}

	potentialStartingPoints := make([]util.Point, 0)
	for _, matchingColumn := range matchingColumns {
		// Y is always greater than X
		if matchingColumn.X+1 == matchingColumn.Y {
			potentialStartingPoints = append(potentialStartingPoints, matchingColumn)
		}
	}

	for _, startingPoint := range potentialStartingPoints {
		leftPointer := startingPoint.X
		rightPointer := startingPoint.Y
		// until one of the pointers is out of bounds
		for leftPointer >= 0 && rightPointer < patternWidth {
			if getColumn(pattern, leftPointer) != getColumn(pattern, rightPointer) {
				break
			}
			leftPointer--
			rightPointer++
		}
		if leftPointer < 0 || rightPointer >= patternWidth {
			return true, startingPoint
		}
	}
	return false, util.Point{}
}

func getColumn(pattern []string, index int) string {
	column := ""
	for _, row := range pattern {
		column += string(row[index])
	}
	return column
}
