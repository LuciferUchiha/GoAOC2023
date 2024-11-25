package main

import (
	"GoAOC2023/util"
	"fmt"
	"log"
)

func main() {
	lines := util.ReadLines("day10/day10.in")

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	// find starting position
	startX := 0
	startY := 0
	for y, row := range grid {
		for x, char := range row {
			if char == 'S' {
				fmt.Printf("Found starting position at %d,%d\n", x, y)
				startX = x
				startY = y
				break
			}
		}
	}

	startingPipe, prevX, prevY := getStartingPipe(grid, startX, startY)
	fmt.Printf("Starting pipe is %c\n", startingPipe)
	grid[startY][startX] = startingPipe
	backAtStart := false
	currentX := startX
	currentY := startY
	steps := 0
	for !backAtStart {
		nextX, nextY := getNextPipe(grid, currentX, currentY, prevX, prevY)
		steps++
		if nextX == startX && nextY == startY {
			backAtStart = true
		}
		prevX = currentX
		prevY = currentY
		currentX = nextX
		currentY = nextY
	}
	fmt.Printf("Walked %d steps\n", steps)
	fmt.Printf("Therefore the answer is %d\n", steps/2)
}

func getStartingPipe(grid [][]rune, x, y int) (rune, int, int) {
	// what piece is at the starting position?
	surroundingLookup := [][]int{
		[]int{0, -1}, // above
		[]int{1, 0},  // right
		[]int{0, 1},  // below
		[]int{-1, 0}, // left
	}
	connectablePipes := [][]rune{
		[]rune{'|', '7', 'F'}, // from above
		[]rune{'-', '7', 'J'}, // from right
		[]rune{'|', 'L', 'J'}, // from below
		[]rune{'-', 'L', 'F'}, // from left
	}

	above := 0
	right := 1
	below := 2
	left := 3

	// two of the surrounding pieces need to be connected to the starting piece
	connectingPipes := [2]int{} // id in lookup table, i.e. 0 = above, 1 = right, 2 = below, 3 = left

	foundPipes := 0
	for i, surrounding := range surroundingLookup {
		lookingAtX := x + surrounding[0]
		lookingAtY := y + surrounding[1]
		// check if we're out of bounds
		if lookingAtX < 0 || lookingAtX >= len(grid[0]) || lookingAtY < 0 || lookingAtY >= len(grid) {
			continue
		}
		// check if the piece is connectable
		for _, connectablePipe := range connectablePipes[i] {
			if grid[lookingAtY][lookingAtX] == connectablePipe {
				connectingPipes[foundPipes] = i
				foundPipes++
				break
			}
		}
	}
	fmt.Println(connectingPipes)
	if foundPipes != 2 {
		log.Fatalf("Found %d connecting pipes, expected 2", foundPipes)
	}

	// now we need to find the starting pipe
	if connectingPipes[0] == above && connectingPipes[1] == right || connectingPipes[1] == above && connectingPipes[0] == right {
		return 'L', surroundingLookup[above][0], surroundingLookup[above][1]
	} else if connectingPipes[0] == right && connectingPipes[1] == below || connectingPipes[1] == right && connectingPipes[0] == below {
		return 'F', surroundingLookup[right][0], surroundingLookup[right][1]
	} else if connectingPipes[0] == below && connectingPipes[1] == left || connectingPipes[1] == below && connectingPipes[0] == left {
		return '7', surroundingLookup[below][0], surroundingLookup[below][1]
	} else if connectingPipes[0] == left && connectingPipes[1] == above || connectingPipes[1] == left && connectingPipes[0] == above {
		return 'J', surroundingLookup[left][0], surroundingLookup[left][1]
	} else if connectingPipes[0] == above && connectingPipes[1] == below || connectingPipes[1] == above && connectingPipes[0] == below {
		return '|', surroundingLookup[above][0], surroundingLookup[above][1]
	} else if connectingPipes[0] == left && connectingPipes[1] == right || connectingPipes[1] == left && connectingPipes[0] == right {
		return '-', surroundingLookup[right][0], surroundingLookup[right][1]
	}
	log.Fatalf("Could not find starting pipe")
	return '.', 0, 0
}

func getNextPipe(grid [][]rune, x, y int, prevX, prevY int) (int, int) {
	char := grid[y][x]
	switch char {
	case '|': // vertical pipe
		// am I going up or down?
		if y > prevY {
			return x, y + 1 // going down
		} else {
			return x, y - 1 // going up
		}
	case '-': // horizontal pipe
		// am I going left or right?
		if x > prevX {
			return x + 1, y // going right
		} else {
			return x - 1, y // going left
		}
	case 'L': // 90 degree bend connecting North and East
		// am I going up or right?
		if y > prevY {
			return x + 1, y // going right
		} else {
			return x, y - 1 // going up
		}
	case 'J': // 90 degree bend connecting North and West
		// am I going up or left?
		if y > prevY {
			return x - 1, y // going left
		} else {
			return x, y - 1 // going up
		}
	case '7': // 90 degree bend connecting South and West
		// am I going down or left?
		if y < prevY {
			return x - 1, y // going left
		} else {
			return x, y + 1 // going down
		}
	case 'F': // 90 degree bend connecting South and East
		// am I going down or right?
		if y < prevY {
			return x + 1, y // going right
		} else {
			return x, y + 1 // going down
		}
	case 'S': // start
		// we stay where we are
		return x, y
	}
	log.Fatalf("Unknown character %c at %d,%d", char, x, y)
	return 0, 0
}
