package main

import (
	"GoAOC2023/util"
	"fmt"
	"slices"
	"strings"
)

type Module struct {
	label     string
	operation string
	neighbors []string
}

type Signal struct {
	current string
	input   int
	source  string
}

const (
	Button      = "button"
	Broadcaster = "broadcaster"
	FlipFlop    = "%"
	Conjunction = "&"
	Deadend     = "deadend"
)

const (
	Low  = 0
	High = 1
)

const (
	Off = 0
	On  = 1
)

func main() {
	lines := util.ReadLines("day20/day20.in")
	modules := make(map[string]Module)
	flipFlopStates := make(map[string]int)
	lastSignalReceived := make(map[string]map[string]int)
	for _, line := range lines {
		tokens := strings.Split(line, " -> ")
		source := tokens[0]
		neighbors := strings.Split(tokens[1], ", ")
		if strings.Contains(source, Broadcaster) {
			modules[Broadcaster] = Module{label: Broadcaster, operation: Broadcaster, neighbors: neighbors}
		} else if strings.Contains(source, FlipFlop) {
			modules[source[1:]] = Module{label: source[1:], operation: FlipFlop, neighbors: neighbors}
			flipFlopStates[source[1:]] = Off
		} else if strings.Contains(source, Conjunction) {
			modules[source[1:]] = Module{label: source[1:], operation: Conjunction, neighbors: neighbors}
			lastSignalReceived[source[1:]] = make(map[string]int)
		} else {
			panic("Unknown Module")
		}
	}

	// for each conjunction, add the Modules have it as a neighbor to the lastSignalReceived map
	for _, module := range modules {
		if module.operation == Conjunction {
			for _, otherModule := range modules {
				if otherModule.label != module.label {
					if slices.Contains(otherModule.neighbors, module.label) {
						lastSignalReceived[module.label][otherModule.label] = Low
					}
				}
			}
		}
	}

	for _, module := range modules {
		fmt.Println(module)
	}

	buttonPresses := 1_000_000_000
	lowsSent := 0
	highsSent := 0
	for i := 0; i < buttonPresses; i++ {
		queue := make([]Signal, 0)
		//println("Button press:", i)
		signalsToMachine := 0
		queue = append(queue, Signal{current: Broadcaster, input: Low, source: Button})
		for len(queue) > 0 {
			signal := queue[0]
			queue = queue[1:]

			input := signal.input
			if input == Low {
				lowsSent++
			} else {
				highsSent++
			}

			currentLabel := signal.current
			sourceLabel := signal.source
			current, currentOk := modules[currentLabel]
			operation := current.operation

			if currentLabel == "rx" && input == Low {
				signalsToMachine++
			}

			if !currentOk {
				operation = Deadend
			}

			output := Low
			//println("Processing:", sourceLabel, "->", currentLabel, "input:", input)
			if operation == Broadcaster {
				output = input
			} else if operation == FlipFlop {
				if input == High {
					goto doNothing
				} else {
					// switch state
					if flipFlopStates[currentLabel] == Off {
						flipFlopStates[currentLabel] = On
						output = High
					} else {
						flipFlopStates[currentLabel] = Off
						output = Low
					}
				}
			} else if operation == Conjunction {
				lastSignalReceived[currentLabel][sourceLabel] = input
				inputs := lastSignalReceived[currentLabel]
				allHigh := true
				for _, value := range inputs {
					if value == Low {
						allHigh = false
						break
					}
				}
				if allHigh {
					output = Low
				} else {
					output = High
				}
			} else if operation == Deadend {
				goto doNothing
			} else {
				println("Unknown operation:", operation)
				panic("Unknown operation")
			}

			// send the message to all neighbors
			for _, neighbor := range current.neighbors {
				queue = append(queue, Signal{current: neighbor, input: output, source: currentLabel})
			}
		doNothing:
		}
		if signalsToMachine == 1 {
			fmt.Println("Machine received 1 signal after", i+1, "button presses")
			break
		}
	}
	fmt.Println(lowsSent, highsSent, lowsSent*highsSent)
}
