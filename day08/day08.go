package main

import (
	"GoAOC2023/util"
	"fmt"
	"regexp"
)

func main() {
	lines := util.ReadLines("day08/day08.in")

	firstLine := lines[0]
	instructions := []rune(firstLine)
	graph := make(map[string]Destinations)

	for i := 2; i < len(lines); i++ {
		line := lines[i]
		fmt.Println(line)
		re := regexp.MustCompile(`(\w+)`)
		places := re.FindAllString(line, -1)
		graph[places[0]] = Destinations{places[1], places[2]}
	}

	currentPlace := "AAA"
	target := "ZZZ"
	currentInstruction := 0
	steps := 0
	for currentPlace != target {
		destinations := graph[currentPlace]
		// we always cycle through the instructions
		if currentInstruction >= len(instructions) {
			currentInstruction = 0
		}
		instruction := instructions[currentInstruction]
		if instruction == 'L' {
			currentPlace = destinations.left
		} else if instruction == 'R' {
			currentPlace = destinations.right
		} else {
			panic("Invalid instruction")
		}
		currentInstruction++
		steps++
	}
	fmt.Println(steps)
}

type Destinations struct {
	left  string
	right string
}
