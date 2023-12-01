package main

import (
	"GoAOC2023/util"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func main() {
	lines := util.ReadLines("day01/day01.in")
	sum := 0
	for _, line := range lines {
		// find first and second number in line and concatenate them to a new number
		re := regexp.MustCompile("[0-9]")
		digits := re.FindAllString(line, -1)
		first := digits[0]
		last := digits[len(digits)-1]
		concatNumber := first + last
		number, err := strconv.Atoi(concatNumber)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	fmt.Println(sum)
}
