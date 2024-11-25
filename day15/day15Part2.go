package main

import (
	"GoAOC2023/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadLines("day15/day15.in")
	tokens := strings.Split(lines[0], ",")
	boxes := make(map[int][]string)
	dict := make(map[string]int)
	sum := 0
	for _, token := range tokens {
		isAdd := true
		operationIndex := strings.LastIndex(token, "=")
		if operationIndex == -1 {
			operationIndex = strings.LastIndex(token, "-")
			isAdd = false
		}
		label := token[:operationIndex]
		if dict[label] == 0 {
			dict[label] = getValue(label)
		}
		focalStrength := token[operationIndex+1:]
		boxIndex := dict[label]
		lens := label + " " + focalStrength

		println("Before")
		fmt.Println(boxIndex, strings.Join(boxes[boxIndex], ", "))

		if isAdd {
			slice := boxes[boxIndex]
			index := -1
			for i := 0; i < len(slice); i++ {
				sliceLabel := strings.Split(slice[i], " ")[0]
				if sliceLabel == label {
					index = i
					break
				}
			}
			if index != -1 {
				println("Setting", lens, "at index", index)
				slice[index] = lens
			} else {
				println("Appending", lens)
				slice = append(slice, lens)
			}
			boxes[boxIndex] = slice
		} else {
			slice := boxes[boxIndex]
			// remove lens from slice
			index := -1
			for i := 0; i < len(slice); i++ {
				sliceLabel := strings.Split(slice[i], " ")[0]
				if sliceLabel == label {
					index = i
					break
				}
			}
			println("Removing", lens, "at index", index)
			if index != -1 {
				slice = slices.Delete(slice, index, index+1)
			}
			boxes[boxIndex] = slice
		}
		println("After")
		fmt.Println(boxIndex, strings.Join(boxes[boxIndex], ", "))
	}

	for i := 0; i < 256; i++ {
		if len(boxes[i]) == 0 {
			continue
		}
		fmt.Println(i, strings.Join(boxes[i], ", "))
		for j, lens := range boxes[i] {
			focalToken := strings.Split(lens, " ")[1]
			focalStrength, err := strconv.Atoi(focalToken)
			if err != nil {
				panic(err)
			}
			sum += (i + 1) * focalStrength * (j + 1)
		}
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
