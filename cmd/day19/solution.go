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

type blueprint struct {
	id                                                                                                                        int
	oreRobotCostOre, clayRobotCostOre, obsidianRobotCostOre, obsidianRobotCostClay, geodeRobotCostOre, geodeRobotCostObsidian int
}

type state struct {
	minute    int
	resources Resources
	robots    Robots
}

type Resources struct {
	ore, clay, obsidian, geode int
}

type Robots struct {
	ore, clay, obsidian, geode int
}

func loadInput(filename string) (blueprints []blueprint) {
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)

	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		r := regexp.MustCompile("\\d+")
		s := r.FindAllString(line, -1)
		blueprints = append(blueprints, blueprint{
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

func whatRobotsCanBeBuilt(x Resources, b blueprint) (robotsToBeBuilt []Robots) {

	robotsSet := make(map[Robots]bool)
	var selected Robots
	findRobotCombinations(x, selected, &robotsSet, b) // side effects to robotsSet

	// extract keys
	keys := make([]Robots, len(robotsSet)+1)
	i := 0
	for k := range robotsSet {
		keys[i] = k
		i++
	}
	keys = append(keys, Robots{})
	//fmt.Println(x, keys)
	return keys
}

func findRobotCombinations(x Resources, selected Robots, robotSet *map[Robots]bool, b blueprint) {
	if _, visited := (*robotSet)[selected]; visited {
		fmt.Println("already visited", selected)
		return
	}

	// try to make a clay robot
	if x.ore/b.clayRobotCostOre > 0 {
		//fmt.Println("can make clay")
		x2 := x                      // copy state
		selectedClay := selected     // copy state
		selectedClay.clay += 1       // add robot
		x2.ore -= b.clayRobotCostOre // subtract robot cost
		(*robotSet)[selectedClay] = true
		//fmt.Println(robotSet)
		findRobotCombinations(x2, selectedClay, robotSet, b)
	}

	if x.ore/b.obsidianRobotCostOre > 0 && x.clay/b.obsidianRobotCostClay > 0 {
		//fmt.Println("can make obsidian")
		x2 := x
		selectedObsidian := selected
		selectedObsidian.obsidian += 1
		x2.ore -= b.obsidianRobotCostOre
		x2.clay -= b.obsidianRobotCostClay
		(*robotSet)[selectedObsidian] = true
		findRobotCombinations(x2, selectedObsidian, robotSet, b)
	}

	//// either build a clay, or obsidian, or geode, or do nothing
	if x.ore/b.geodeRobotCostOre > 0 && x.obsidian/b.geodeRobotCostObsidian > 0 {
		//fmt.Println("can make geode")
		x2 := x
		selectedGeode := selected
		selectedGeode.geode += 1
		x2.ore -= b.geodeRobotCostOre
		x2.obsidian -= b.geodeRobotCostObsidian
		(*robotSet)[selectedGeode] = true
		findRobotCombinations(x2, selectedGeode, robotSet, b)
	}
}

func mineResources(x state) state {
	x.resources.ore += x.robots.ore
	x.resources.clay += x.robots.clay
	x.resources.obsidian += x.robots.obsidian
	x.resources.geode += x.robots.geode
	return x
}

func addRobots(x state, robots Robots) state {
	x.robots.ore += robots.ore
	x.robots.clay += robots.clay
	x.robots.obsidian += robots.obsidian
	x.robots.geode += robots.geode
	return x
}

func simulate(x state, b blueprint, dp *map[state]int) (nGeodes int) {
	if nGeodes, visited := (*dp)[x]; visited {
		return nGeodes
	}

	fmt.Printf("Minute %d\n", x.minute)
	if x.minute == 5 {
		return x.resources.geode
	}
	// what action can we take?
	actions := whatRobotsCanBeBuilt(x.resources, b)

	maxFinalGeodes := 0
	for _, action := range actions {
		x2 := oneMinute(x, action)
		fmt.Println(x2)
		nGeodes = simulate(x2, b, dp)
		if nGeodes > maxFinalGeodes {
			maxFinalGeodes = nGeodes
		}
	}
	return maxFinalGeodes
}

func oneMinute(x state, robotsToAdd Robots) state {
	// perform mining
	x = mineResources(x)

	// add robots
	x = addRobots(x, robotsToAdd)

	// increment time
	x.minute += 1
	return x
}

func part1(filename string) int {
	blueprints := loadInput(filename)
	fmt.Println("blueprint, ", blueprints)
	totalQualityLevel := 0
	for _, blueprint := range blueprints {
		x := state{
			0,
			Resources{0, 0, 0, 0},
			Robots{1, 0, 0, 0},
		}
		dp := make(map[state]int)
		nGeodes := simulate(x, blueprint, &dp)
		totalQualityLevel += nGeodes * blueprint.id
		fmt.Println("total quality level: ", totalQualityLevel)
	}

	return totalQualityLevel
}

func part2(filename string) int {
	return 0
}
