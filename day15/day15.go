package main

import (
	"GoAOC2023/util"
	"strings"
)

func main() {
	lines := util.ReadLines("day15/day15.in")
	tokens := strings.Split(lines[0], ",")
	dict := make(map[string]int)
	sum := 0
	for _, token := range tokens {
		if dict[token] == 0 {
			dict[token] = getValue(token)
		}
		sum += dict[token]
	}
	println(sum)
}

func getValue(token string) int {
	value := 0
	for _, char := range token {
		value += int(char)
		value *= 17
		value %= 256
	}
	return value
}
