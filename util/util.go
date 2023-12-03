package util

import (
	"bufio"
	"log"
	"os"
	"regexp"
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
