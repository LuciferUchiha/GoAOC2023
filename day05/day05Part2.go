package main

import (
	"GoAOC2023/util"
	"fmt"
	"log"
	"math"
)

func main() {
	lines := util.ReadLines("day05/day05.in")

	// extract the seeds
	currentLine := 0
	seedsRanges := util.ExtractNumbers(lines[currentLine])
	currentLine += 3

	fmt.Println("We have ", len(seedsRanges)/2, " seed ranges")
	fmt.Println("Seed ranges: ", seedsRanges)

	// generate all maps
	seedToSoilMap := generateMap(lines, currentLine)
	currentLine += len(seedToSoilMap) + 2 // skip the empty line and the header
	soilToFertilizerMap := generateMap(lines, currentLine)
	currentLine += len(soilToFertilizerMap) + 2
	fertilizerToWaterMap := generateMap(lines, currentLine)
	currentLine += len(fertilizerToWaterMap) + 2
	waterToLightMap := generateMap(lines, currentLine)
	currentLine += len(waterToLightMap) + 2
	lightToTemperatureMap := generateMap(lines, currentLine)
	currentLine += len(lightToTemperatureMap) + 2
	temperatureToHumidityMap := generateMap(lines, currentLine)
	currentLine += len(temperatureToHumidityMap) + 2
	humidityToLocationMap := generateMap(lines, currentLine)
	currentLine += len(humidityToLocationMap) + 2

	minLocation := math.MaxInt

	for i := 0; i < len(seedsRanges); i += 2 {
		// for each seed in the range
		seedStart := seedsRanges[i]
		length := seedsRanges[i+1]
		for seed := seedStart; seed < seedStart+length; seed++ {
			soil := passValueThroughMap(seed, seedToSoilMap)
			fertilizer := passValueThroughMap(soil, soilToFertilizerMap)
			water := passValueThroughMap(fertilizer, fertilizerToWaterMap)
			light := passValueThroughMap(water, waterToLightMap)
			temperature := passValueThroughMap(light, lightToTemperatureMap)
			humidity := passValueThroughMap(temperature, temperatureToHumidityMap)
			location := passValueThroughMap(humidity, humidityToLocationMap)
			if location < minLocation {
				minLocation = location
			}
		}
	}
	fmt.Println("The minimum location is: ", minLocation)
}

func generateMap(lines []string, currentLine int) [][]int {
	sourceTargetMap := make([][]int, 0)
	for currentLine < len(lines) && lines[currentLine] != "" {
		tokens := util.ExtractNumbers(lines[currentLine])
		if len(tokens) != 3 {
			log.Fatal("Invalid number of tokens")
		}
		sourceTargetMap = append(sourceTargetMap, tokens)
		currentLine++
	}
	return sourceTargetMap
}

func passValueThroughMap(source int, sourceTargetMap [][]int) int {
	for _, sourceTargetPair := range sourceTargetMap {
		targetStart := sourceTargetPair[0]
		sourceStart := sourceTargetPair[1]
		length := sourceTargetPair[2]
		if source >= sourceStart && source <= sourceStart+length {
			return targetStart + (source - sourceStart)
		}
	}
	return source
}
