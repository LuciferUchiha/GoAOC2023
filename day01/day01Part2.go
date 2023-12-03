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
		line = ReplaceWrittenNumbersWithDigits(line)
		re := regexp.MustCompile("[1-9]")
		digits := re.FindAllString(line, -1)

		first := digits[0]
		last := digits[len(digits)-1]

		fmt.Println("first:", first)
		fmt.Println("last:", last)

		concatNumber := first + last

		fmt.Println("concatNumber:", concatNumber)

		number, err := strconv.Atoi(concatNumber)
		if err != nil {
			log.Fatal(err)
		}

		sum += number
	}
	fmt.Println(sum)
}

func ReplaceWrittenNumbersWithDigits(line string) string {
	for true {
		newLine := line
		re := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine")
		writtenNumbersIndexes := re.FindAllStringIndex(line, -1)
		if len(writtenNumbersIndexes) == 0 {
			break
		}
		for _, writtenNumberIndex := range writtenNumbersIndexes {
			writtenNumber := line[writtenNumberIndex[0]:writtenNumberIndex[1]]
			replacementNumber := GetReplacementNumber(writtenNumber)
			newLine = line[:writtenNumberIndex[0]] + replacementNumber + line[writtenNumberIndex[1]:]
		}
		line = newLine
	}
	return line
}

func GetReplacementNumber(writtenNumber string) string {
	switch writtenNumber {
	case "one":
		return "o1e"
	case "two":
		return "t2o"
	case "three":
		return "t3hree"
	case "four":
		return "f4ur"
	case "five":
		return "f5ve"
	case "six":
		return "s6x"
	case "seven":
		return "s7ven"
	case "eight":
		return "e8ght"
	case "nine":
		return "n9ne"
	default:
		log.Fatal("Unknown written number:", writtenNumber)
	}
	return "0"
}
