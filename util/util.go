package util

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

func ManhattanDistance(point1, point2 Point) float64 {
	return math.Abs(float64(point1.X-point2.X)) + math.Abs(float64(point1.Y-point2.Y))
}

func TwoDimRuneSlicesEqual(slice1, slice2 [][]rune) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for y, row := range slice1 {
		if len(row) != len(slice2[y]) {
			return false
		}
		for x, symbol := range row {
			if symbol != slice2[y][x] {
				return false
			}
		}
	}
	return true
}

type Point struct {
	X int
	Y int
}

func (p Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

type Pair struct {
	A int
	B int
}

type GridLine struct {
	Start Point
	End   Point
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (line GridLine) String() string {
	return fmt.Sprintf("%s -> %s", line.Start, line.End)
}
