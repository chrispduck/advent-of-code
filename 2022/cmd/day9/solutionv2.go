package main

import (
	"advent-of-code/cmd/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	Part2NKnots = 10
)

func main() {
	fmt.Println(part1("example_input.txt"))
	fmt.Println(part1("input.txt"))
	fmt.Println(part2("example_input_part2.txt"))
	fmt.Println(part2("input.txt"))
}

type command struct {
	direction string
	steps     int
}

func loadInput(filename string) []command {
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)
	var commands []command
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		steps := utils.StrToInt(line[2:])
		commands = append(commands, command{
			direction: string(direction),
			steps:     steps,
		})
	}
	utils.CheckErr(scanner.Err())
	return commands
}

// idempotent
func updateVisited(c coordinate, visited *map[string]int) {
	key := "(" + strconv.Itoa(c.x) + "," + strconv.Itoa(c.y) + ")"
	(*visited)[key] += 1
}

type coordinate struct {
	x int
	y int
}

func (c1 coordinate) subtract(c2 coordinate) coordinate {
	return coordinate{
		x: c1.x - c2.x,
		y: c1.y - c2.y,
	}
}

func (c1 coordinate) add(c2 coordinate) coordinate {
	return coordinate{
		x: c1.x + c2.x,
		y: c1.y + c2.y,
	}
}

func (c1 coordinate) elfDistance(c2 coordinate) int {
	return utils.Max(utils.Abs(c1.y-c2.y), utils.Abs(c1.x-c2.x))

}

func moveOne(c coordinate, direction string) coordinate {
	switch direction {
	case "U":
		return c.add(coordinate{0, 1})
	case "D":
		return c.add(coordinate{0, -1})
	case "L":
		return c.add(coordinate{-1, 0})
	case "R":
		return c.add(coordinate{1, 0})
	default:
		log.Fatalln("failed to parse command")
	}
	return coordinate{}
}

func moveIfRequired(c, cH coordinate) (res coordinate) {
	if c.elfDistance(cH) <= 1 {
		// no move
		return c
	}

	// compute deltas
	delta := cH.subtract(c)

	// select move
	var move coordinate
	if delta.x > 0 {
		move.x = 1
	} else if delta.x < 0 {
		move.x = -1
	}
	if delta.y > 0 {
		move.y = 1
	} else if delta.y < 0 {
		move.y = -1
	}

	// execute move
	return c.add(move)
}

func part1(filename string) int {
	commands := loadInput(filename)

	// initialise
	var head, tail coordinate
	visited := make(map[string]int)
	updateVisited(tail, &visited)

	for _, cmd := range commands {
		for cmd.steps > 0 {
			// move head
			head = moveOne(head, cmd.direction)
			cmd.steps -= 1
			// does tail need to move?
			tail = moveIfRequired(tail, head)
			updateVisited(tail, &visited)
		}
	}
	return len(visited)
}

func part2(filename string) int {
	commands := loadInput(filename)
	// initialise
	rope := make([]coordinate, Part2NKnots)
	visited := make(map[string]int)
	tail := &rope[len(rope)-1]
	updateVisited(*tail, &visited)

	for _, cmd := range commands {
		for cmd.steps > 0 {
			// move head
			rope[0] = moveOne(rope[0], cmd.direction)
			cmd.steps -= 1

			// for each part of the rope chain
			for i := 1; i < len(rope); i++ {
				rope[i] = moveIfRequired(rope[i], rope[i-1])
			}
			updateVisited(*tail, &visited)
		}
	}
	return len(visited)
}
