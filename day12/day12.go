package main

import (
	"GoAOC2023/util"
	"fmt"
	"strings"
)

const Dot = '.'
const Hash = '#'
const Unknown = '?'

func main() {
	lines := util.ReadLines("day12/day12.in")
	rows := make([]string, len(lines))
	constraints := make([][]int, len(lines))
	for i, line := range lines {
		tokens := strings.Split(line, " ")
		rows[i] = strings.TrimSpace(tokens[0])
		constraints[i] = util.ExtractNumbers(tokens[1])
	}

	sum := 0
	for i, row := range rows {
		fmt.Println("Row", i)
		fmt.Println(row, constraints[i])
		res := numPossibleRows(row, constraints[i])
		fmt.Println(res)
		sum += res
	}
	fmt.Println(sum)
}

func numPossibleRows(row string, constraints []int) int {
	// Did we run out of constraints?
	if len(constraints) == 0 {
		// if there are no more hashes, then we're valid
		if !strings.Contains(row, string(Hash)) {
			return 1
		} else {
			// still have hashes, we're invalid
			return 0
		}
	}

	// There are more constraints, but no more spaces
	if len(row) == 0 {
		return 0
	}

	// Look at the next element
	nextCharacter := row[0]
	nextGroup := constraints[0]

	pound := func() int {
		if nextGroup > len(row) {
			// We can't fit the next group, abort
			return 0
		}

		// If the first is a hash, then the first n characters must be hashes
		thisGroup := row[:nextGroup]
		thisGroup = strings.ReplaceAll(thisGroup, "?", "#")

		// If we can't fit the constraint
		if thisGroup != strings.Repeat("#", nextGroup) {
			return 0
		}

		// If the rest is just the last constaint
		if len(row) == nextGroup {
			// Make sure this is the last group
			if len(constraints) == 1 {
				return 1
			} else {
				// There's more groups, we can't make it work
				return 0
			}
		}

		// Make sure the character can be a separator
		if row[nextGroup] == '?' || row[nextGroup] == '.' {
			return numPossibleRows(row[nextGroup+1:], constraints[1:])
		}

		// Can't be handled, there are no possibilities
		return 0
	}

	dot := func() int {
		// We just skip over the dot looking for the next hash
		return numPossibleRows(row[1:], constraints)
	}

	var out int
	if nextCharacter == Hash {
		out = pound()
	} else if nextCharacter == Dot {
		out = dot()
	} else if nextCharacter == Unknown {
		// This character could be either character, so we'll explore both possibilities
		out = dot() + pound()
	} else {
		panic("Unknown character")
	}
	return out
}
