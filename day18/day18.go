package main

import (
	"GoAOC2023/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadLines("day18/day18.in")
	directions := make([]rune, len(lines))
	steps := make([]int, len(lines))
	colors := make([]string, len(lines))

	for i, line := range lines {
		tokens := strings.Split(line, " ")
		directions[i] = rune(tokens[0][0])
		number, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}
		steps[i] = number
		colors[i] = tokens[2]
	}

	dir := make(map[rune]util.Point)
	dir['R'] = util.Point{1, 0}
	dir['L'] = util.Point{-1, 0}
	dir['U'] = util.Point{0, -1}
	dir['D'] = util.Point{0, 1}

	points := make([]util.Point, 0)
	prevPoint := util.Point{0, 0}
	borderLength := 1
	points = append(points, prevPoint)
	for i, direction := range directions {
		nextPoint := nextPos(prevPoint, dir[direction], steps[i])
		points = append(points, nextPoint)
		borderLength += steps[i]
		prevPoint = nextPoint
	}

	area := 0
	for i := 0; i < len(points)-1; i++ {
		area += (points[i].Y + points[i+1].Y) * (points[i].X - points[i+1].X)
	}
	area = int(math.Abs(float64(area)))
	fmt.Println("Area:", (area+borderLength+1)/2)
}

func nextPos(a, b util.Point, factor int) util.Point {
	return util.Point{a.X + factor*b.X, a.Y + factor*b.Y}
}