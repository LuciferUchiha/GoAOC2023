package util

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadLines(filename string) []string {
	linesSlice := make([]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		linesSlice = append(linesSlice, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return linesSlice
}

func IsNumber(symbol rune) bool {
	re := regexp.MustCompile("[0-9]")
	return re.MatchString(string(symbol))
}

func ArrayContains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func StringArrayToIntArray(stringArray []string) []int {
	intArray := make([]int, len(stringArray))
	for i, stringNumber := range stringArray {
		number, err := strconv.Atoi(stringNumber)
		if err != nil {
			log.Fatal(err)
		}
		intArray[i] = number
	}
	return intArray
}

func StringArrayRemoveEmptyStrings(stringArray []string) []string {
	newStringArray := make([]string, 0)
	for _, string := range stringArray {
		if string != "" {
			newStringArray = append(newStringArray, string)
		}
	}
	return newStringArray
}

func StringArrayTrimElements(stringArray []string) []string {
	newStringArray := make([]string, 0)
	for _, text := range stringArray {
		newStringArray = append(newStringArray, strings.TrimSpace(text))
	}
	return newStringArray
}

func ExtractNumber(text string) int {
	re := regexp.MustCompile("-*\\w+")
	numberString := re.FindString(text)
	number, err := strconv.Atoi(numberString)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

func ExtractNumbers(text string) []int {
	re := regexp.MustCompile("-*\\w+")
	numberStrings := re.FindAllString(text, -1)
	numbers := make([]int, 0)
	for _, numberString := range numberStrings {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func ConcatNumbersInIntArr(numbers []int) int {
	numberString := ""
	for _, number := range numbers {
		numberString += strconv.Itoa(number)
	}
	number, err := strconv.Atoi(numberString)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

type Pair[T, U any] struct {
	First  T
	Second U
}
