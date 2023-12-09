package main

import (
	"GoAOC2023/util"
	"fmt"
)

func main() {
	lines := util.ReadLines("day09/day09.in")

	rootSequences := make([][]int, len(lines))
	for i, line := range lines {
		rootSequences[i] = util.ExtractNumbers(line)
	}

	predictionSum := 0

	for _, rootSequence := range rootSequences {
		sequences := make([][]int, 0)
		sequences = append(sequences, rootSequence)
		sequenceIndex := len(sequences) - 1
		fmt.Println(rootSequence)
		for !areAllZero(sequences[sequenceIndex]) {
			parentSequence := sequences[sequenceIndex]
			length := len(parentSequence) - 1
			newSequence := make([]int, length)

			// the newSequence contains the difference between the second item and the first item etc.
			for i := 0; i < length; i++ {
				newSequence[i] = parentSequence[i+1] - parentSequence[i]
			}
			fmt.Println(newSequence)
			sequences = append(sequences, newSequence)
			sequenceIndex++
		}

		// to predict the next number, we need to stat with the last element of the last subsequence
		// and then add it to the last element of the parent sequence and so on
		prediction := 0
		// we can do one less iteration because the last sequence is all zeros
		for i := len(sequences) - 2; i >= 0; i-- {
			parentSequence := sequences[i]
			prevPrediction := prediction
			prediction = parentSequence[0] - prediction
			fmt.Println(prediction, " = ", parentSequence[0], " - ", prevPrediction)
		}
		fmt.Println(prediction)
		predictionSum += prediction
	}
	fmt.Println(predictionSum)
}

func areAllZero(sequence []int) bool {
	for _, val := range sequence {
		if val != 0 {
			return false
		}
	}
	return true
}
