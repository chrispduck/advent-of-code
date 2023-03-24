package main

import (
	"advent-of-code/cmd/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	//fmt.Println(part1("example_input.txt"))
	fmt.Println(part1("input.txt"))
	//fmt.Println(part2("example_input.txt"))
	//fmt.Println(part2("input.txt"))
}

type State struct {
	minutesLeft int
	score       int
	position    string
	opened      map[string]interface{}
	path        string
}

type ScorePath struct {
	score int
	path  string
}
type Rules struct {
	flowRates map[string]int
	graph     map[string][]string
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
		minutesLeft: 30,
		score:       0,
		position:    "AA",
		opened:      make(map[string]interface{}),
		path:        "A",
	}

	rules := Rules{
		flowRates: flowRates,
		graph:     graph,
	}

	memo := make(map[string]ScorePath)
	res := findMaxPressureRelief(state, rules, memo)
	return res
}

func makeKey(state State) string {
	return fmt.Sprintf("%d-%v-%v-%d", state.minutesLeft, state.opened, state.position, state.score)
}

func findMaxPressureRelief(state State, rules Rules, memo map[string]ScorePath) int {
	//fmt.Println(state)
	if state.minutesLeft <= 0 {
		//fmt.Println("finished", state.score, state.opened, state.path)
		return state.score
	}

	//is the path worth exploring?
	key := makeKey(state)
	//fmt.Println(key)
	if val, exists := memo[key]; exists {
		return val.score
	}

	// explore the path
	state.score += releasePressureOneMinute(state.opened, rules.flowRates)
	state.minutesLeft--

	// find the best path from here
	bestScore := state.score

	// if current valve isn't open, open it
	//fmt.Println(state.position, state.opened)
	if _, isOpen := state.opened[state.position]; !isOpen && rules.flowRates[state.position] > 0 {
		_state := copyState(state)
		_state.opened[state.position] = true
		_state.path += "+"
		_score := findMaxPressureRelief(_state, rules, memo)
		bestScore = utils.Max(_score, bestScore)
	}

	// or follow any of the paths
	for _, to := range rules.graph[state.position] {
		_state := copyState(state)
		_state.position = to
		_state.path += string(to[0])
		_score := findMaxPressureRelief(_state, rules, memo)
		bestScore = utils.Max(_score, bestScore)
	}

	//if v := memo[key].score; v < bestScore {
	memo[key] = ScorePath{
		score: bestScore,
		path:  state.path,
	}
	//}
	return bestScore
}

func releasePressureOneMinute(opened map[string]interface{}, flowRates map[string]int) int {
	p := 0
	for valve, _ := range opened {
		p += flowRates[valve]
	}
	return p
}

func copyState(state State) State {
	opened := make(map[string]interface{})
	for k, v := range state.opened {
		opened[k] = v
	}
	return State{
		minutesLeft: state.minutesLeft,
		score:       state.score,
		position:    state.position,
		opened:      opened,
		path:        state.path,
	}
}

func part2(filename string) int {
	flowRates, graph := loadInput(filename)
	fmt.Println(flowRates)
	fmt.Println(graph)
	fmt.Println(flowRates)
	fmt.Println(graph)
	state := State{
		minutesLeft: 26,
		score:       0,
		position:    "AA",
		opened:      make(map[string]interface{}),
		path:        "A",
	}
	rules := Rules{
		flowRates: flowRates,
		graph:     graph,
	}

	memo := make(map[string]ScorePath)
	res := findMaxPressureRelief(state, rules, memo)
	fmt.Println(memo)
	return res
}

func simplifyBidirectionalGraph(nodes map[string][]string, flowRates map[string]int) (nodesAfter map[string][]string, flowRatesAfter map[string]int) {
	var nodesToRemove []string
	for node, _ := range nodes {
		if flowRates[node] == 0 {
			nodesToRemove = append(nodesToRemove, node)
			for _, to := range nodes[node] {
				nodes[to] = utils.IdempotentRemove(nodes[to], node)
				nodes[to] = utils.IdempotentAdds(nodes[to], nodes[node])
			}

		}
	}
	for _, node := range nodesToRemove {
		delete(nodes, node)
		delete(flowRates, node)
	}
	return nodes, flowRates
}
