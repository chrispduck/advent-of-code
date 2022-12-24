package main

import (
	"advent-of-code/cmd/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	fmt.Println(part1("example_input.txt"))
	//fmt.Println(part1("input.txt"))
	//fmt.Println(part2("example_input.txt"))
	//fmt.Println(part2("input.txt"))
}

type State struct {
	minute   int
	score    int
	position string
	opened   map[string]interface{}
}

type Rules struct {
	flowRates map[string]int
	graph     map[string][]string
}

// returns true if A is a superset of B.
//ie if A union B >= B
func isSuperset(a map[string]interface{}, b map[string]interface{}) bool {
	for key, _ := range b {
		_, exists := a[key]
		if !exists {
			return false
		}
	}
	return true
}

func loadInput(filename string) (flowRates map[string]int, graph map[string][]string) {
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)
	flowRates = make(map[string]int)
	graph = make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		rNumber := regexp.MustCompile("\\d+")
		rValves := regexp.MustCompile("[A-Z][A-Z]")
		flowRate := utils.StrToInt(rNumber.FindAllString(line, -1)[0])
		valveNames := rValves.FindAllString(line, -1)
		flowRates[valveNames[0]] = flowRate
		graph[valveNames[0]] = valveNames[1:]
	}
	utils.CheckErr(scanner.Err())
	return flowRates, graph

}

func part1(filename string) int {
	flowRates, graph := loadInput(filename)
	fmt.Println(flowRates)
	fmt.Println(graph)
	state := State{
		minute:   0,
		score:    0,
		position: "AA",
		opened:   make(map[string]interface{}),
	}
	rules := Rules{
		flowRates: flowRates,
		graph:     graph,
	}

	memoBestScore := make(map[string]int)
	res := findMaxPressureRelief(state, rules, memoBestScore)

	return res
}

func makeKey(state State) string {
	return fmt.Sprintf("%d-%s-%s", state.minute, state.position, state.opened)
}

func findMaxPressureRelief(state State, rules Rules, memoBestScore map[string]int) int {
	state.minute += 1
	if state.minute > 30 {
		fmt.Println(state.score, state.opened)
		return state.score
	}

	state.score += releasePressureOneMinute(state.opened, rules.flowRates)

	// is the path worth exploring?
	key := makeKey(state)
	if state.score < memoBestScore[key] {
		return memoBestScore[key]
	}
	memoBestScore[key] = state.score

	// find the best path from here
	var finalState State
	var finalScore int

	// if current valve isn't open, open it
	if _, isOpen := state.opened[state.position]; !isOpen && rules.flowRates[state.position] > 0 {
		_state := state
		_state.opened[state.position] = true
		_score := findMaxPressureRelief(_state, rules, memoBestScore)
		if _score > finalScore {
			finalState = _state
			finalScore = _score
		}
	}

	// or follow any of the paths
	for _, to := range rules.graph[state.position] {
		//fmt.Println("making move from ", state.position, " to ", to, ". ", state)
		_state := state
		_state.position = to
		_score := findMaxPressureRelief(_state, rules, memoBestScore)
		if _score > finalScore {
			finalState = _state
			finalScore = _score
		}
	}
	_ = finalState
	return finalScore
}

func releasePressureOneMinute(opened map[string]interface{}, flowRates map[string]int) int {
	p := 0
	for valve, _ := range opened {
		p += flowRates[valve]
	}
	return p
}

func part2(filename string) int {
	return 0
}
