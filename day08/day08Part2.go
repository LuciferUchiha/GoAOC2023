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

	startingPlaces := make([]string, 0)
	for key, _ := range graph {
		if key[len(key)-1] == 'A' {
			startingPlaces = append(startingPlaces, key)
		}
	}
	fmt.Println(startingPlaces)
	currentPlaces := startingPlaces

	distancesToTarget := make([]int, len(startingPlaces))
	for i, currentPlace := range currentPlaces {
		steps := 0
		for currentPlace[len(currentPlace)-1] != 'Z' {
			instruction := instructions[steps%len(instructions)]
			if instruction == 'L' {
				currentPlace = graph[currentPlace].left
			} else {
				currentPlace = graph[currentPlace].right
			}
			steps++
		}
		distancesToTarget[i] = steps
	}
	fmt.Println(distancesToTarget)
	fmt.Println(LCM(distancesToTarget...))
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(numbers ...int) int {
	if len(numbers) < 2 {
		panic("LCM requires at least two numbers")
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = result * numbers[i] / GCD(result, numbers[i])
	}

	return result
}

type Destinations struct {
	left  string
	right string
}
