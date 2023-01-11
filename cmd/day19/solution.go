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
	fmt.Println(part1("input.txt"))
	fmt.Println(part2("example_input.txt"))
	fmt.Println(part2("input.txt"))
}

var (
	bestScore = 0
	nVisited  = 0
)

type Blueprint struct {
	id                                                                                                                        int
	oreRobotCostOre, clayRobotCostOre, obsidianRobotCostOre, obsidianRobotCostClay, geodeRobotCostOre, geodeRobotCostObsidian int
}

type State struct {
	minutesLeft int
	resources   Resources
	robots      Robots
}

type Resources struct {
	ore, clay, obsidian, geode int
}

type Robots struct {
	ore, clay, obsidian, geode int
}

func loadInput(filename string) (blueprints []Blueprint) {
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)

	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		r := regexp.MustCompile("\\d+")
		s := r.FindAllString(line, -1)
		blueprints = append(blueprints, Blueprint{
			i,
			utils.StrToInt(s[1]),
			utils.StrToInt(s[2]),
			utils.StrToInt(s[3]),
			utils.StrToInt(s[4]),
			utils.StrToInt(s[5]),
			utils.StrToInt(s[6]),
		})
		i++
	}
	utils.CheckErr(scanner.Err())
	return blueprints
}

func findRobotCombinations(x Resources, b Blueprint) []Robots {
	var robots []Robots
	if x.ore/b.geodeRobotCostOre > 0 && x.obsidian/b.geodeRobotCostObsidian > 0 {
		// definitely build a geode robot
		return []Robots{{0, 0, 0, 1}}
	}
	if x.ore/b.obsidianRobotCostOre > 0 && x.clay/b.obsidianRobotCostClay > 0 {
		robots = append(robots, Robots{0, 0, 1, 0})
	}
	if x.ore/b.clayRobotCostOre > 0 {
		robots = append(robots, Robots{0, 1, 0, 0})
	}
	if x.ore/b.oreRobotCostOre > 0 {
		robots = append(robots, Robots{1, 0, 0, 0})
	}
	robots = append(robots, Robots{0, 0, 0, 0})
	//fmt.Println("robots", robots)
	//fmt.Println("resources", x)

	return robots
}

func (x *State) buildRobot(b Blueprint, toBuild Robots) {
	if toBuild.geode > 0 {
		x.resources.ore -= b.geodeRobotCostOre
		x.resources.obsidian -= b.geodeRobotCostObsidian
		x.robots.geode += 1
	}
	if toBuild.obsidian > 0 {
		x.resources.ore -= b.obsidianRobotCostOre
		x.resources.clay -= b.obsidianRobotCostClay
		x.robots.obsidian += 1
	}
	if toBuild.clay > 0 {
		x.resources.ore -= b.clayRobotCostOre
		x.robots.clay += 1
	}
	if toBuild.ore > 0 {
		x.resources.ore -= b.oreRobotCostOre
		x.robots.ore += 1
	}
}

func (x *State) mineResources() {
	x.resources.ore += x.robots.ore
	x.resources.clay += x.robots.clay
	x.resources.obsidian += x.robots.obsidian
	x.resources.geode += x.robots.geode
}

func (x *State) addRobots(robots Robots) {
	x.robots.ore += robots.ore
	x.robots.clay += robots.clay
	x.robots.obsidian += robots.obsidian
	x.robots.geode += robots.geode
}

func (x *State) oneMinute(b Blueprint, robotsToAdd Robots) {
	x.mineResources()
	x.buildRobot(b, robotsToAdd)
	x.minutesLeft -= 1
}

func simulate(x State, b Blueprint, visited map[State]bool) (nGeodes int, finalState State) {
	//fmt.Printf("Minute %d\n", x.minutesLeft)
	maxPossibleGeodes := x.resources.geode + x.robots.geode*x.minutesLeft + max((x.minutesLeft)*(x.minutesLeft-1)/2, 0)
	if x.minutesLeft <= 0 || maxPossibleGeodes < bestScore {
		//if x.minutesLeft <= 0 {
		return x.resources.geode, x
	}

	// memo visited states
	if _, ok := visited[x]; ok {
		nVisited++
		if nVisited%1000000 == 0 {
			//fmt.Println("nVisited", nVisited)
		}
		return 0, x
	} else {
		visited[x] = true
	}

	// what action can we take?
	actions := findRobotCombinations(x.resources, b)

	for _, action := range actions {

		x2 := deepCopy(x)
		x2.oneMinute(b, action)
		//fmt.Println("state x ", x2)
		//fmt.Println(x2.robots.geode)

		nGeodes, x2 = simulate(x2, b, visited)
		if nGeodes > bestScore {
			bestScore = nGeodes
			finalState = x2
			//fmt.Println("new max", maxFinalGeodes, x2)
		}
	}
	return bestScore, finalState
}

func part1(filename string) int {
	blueprints := loadInput(filename)
	//fmt.Println("Blueprint, ", blueprints)
	totalQualityLevel := 0
	for _, blueprint := range blueprints {
		bestScore = 0
		visited := make(map[State]bool)
		x := State{
			24,
			Resources{0, 0, 0, 0},
			Robots{1, 0, 0, 0},
		}
		nGeodes, x := simulate(x, blueprint, visited)
		fmt.Println("Blueprint ", blueprint.id, "nGeodes", nGeodes)
		//fmt.Println("x", x)
		totalQualityLevel += nGeodes * blueprint.id
		fmt.Println("total quality level: ", totalQualityLevel)
	}

	return totalQualityLevel
}

func part2(filename string) int {
	blueprints := loadInput(filename)
	fmt.Println("Blueprint, ", blueprints)
	quadraticTotal := 1
	if len(blueprints) > 3 {
		blueprints = blueprints[:3]
	}
	for _, blueprint := range blueprints {
		bestScore = 0
		visited := make(map[State]bool)
		x := State{
			32,
			Resources{0, 0, 0, 0},
			Robots{1, 0, 0, 0},
		}
		nGeodes, x := simulate(x, blueprint, visited)
		fmt.Println("Blueprint ", blueprint.id, "nGeodes", nGeodes)
		//fmt.Println("x", x)
		quadraticTotal *= nGeodes
		fmt.Println("quadratic total: ", quadraticTotal)
	}

	return quadraticTotal
}

func deepCopy(x State) State {
	return State{
		x.minutesLeft,
		Resources{x.resources.ore, x.resources.clay, x.resources.obsidian, x.resources.geode},
		Robots{x.robots.ore, x.robots.clay, x.robots.obsidian, x.robots.geode},
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
