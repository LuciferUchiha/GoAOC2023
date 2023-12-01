package util

import (
	"bufio"
	"log"
	"os"
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
