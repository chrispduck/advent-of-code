package main

import (
	"advent-of-code/cmd/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(part1("small_input.txt"))
	fmt.Println(part1("example_input.txt"))
	fmt.Println(part1("input.txt"))
	fmt.Println(part1("small_input.txt"))
	fmt.Println(part2("example_input.txt"))
	fmt.Println(part2("input.txt"))
}

type grid struct {
	elves map[utils.Coordinate]bool
}

func (g *grid) isElf(c utils.Coordinate) bool {
	ok, _ := g.elves[c]
	return ok
}
func (g *grid) getDimensions() (xmin, ymin, xmax, ymax int) {
	xmin, ymin, xmax, ymax = 0, 0, 0, 0
	for c := range g.elves {
		if c.X < xmin {
			xmin = c.X
		}
		if c.X > xmax {
			xmax = c.X
		}
		if c.Y < ymin {
			ymin = c.Y
		}
		if c.Y > ymax {
			ymax = c.Y
		}
	}
	return xmin, ymin, xmax, ymax
}

func (g *grid) nEmptyTiles() int {
	xmin, ymin, xmax, ymax := g.getDimensions()
	nElves := len(g.elves)
	return (xmax-xmin+1)*(ymax-ymin+1) - nElves
}
func (g *grid) print() {
	xmin, ymin, xmax, ymax := g.getDimensions()
	for y := ymin; y <= ymax; y++ {
		for x := xmin; x <= xmax; x++ {
			if g.elves[utils.Coordinate{X: x, Y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
func (g *grid) addElf(c utils.Coordinate) {
	g.elves[c] = true
}

func (g *grid) isElfNorth(c utils.Coordinate) bool {
	n := g.elves[utils.Coordinate{X: c.X, Y: c.Y - 1}]
	nw := g.elves[utils.Coordinate{X: c.X - 1, Y: c.Y - 1}]
	ne := g.elves[utils.Coordinate{X: c.X + 1, Y: c.Y - 1}]
	return (n || nw || ne)
}
func (g *grid) isElfSouth(c utils.Coordinate) bool {
	return (g.elves[utils.Coordinate{X: c.X, Y: c.Y + 1}] || g.elves[utils.Coordinate{X: c.X - 1, Y: c.Y + 1}] || g.elves[utils.Coordinate{X: c.X + 1, Y: c.Y + 1}])
}

func (g *grid) isElfEast(c utils.Coordinate) bool {
	return (g.elves[utils.Coordinate{X: c.X + 1, Y: c.Y + 1}] || g.elves[utils.Coordinate{X: c.X + 1, Y: c.Y}] || g.elves[utils.Coordinate{X: c.X + 1, Y: c.Y - 1}])
}

func (g *grid) isElfWest(c utils.Coordinate) bool {
	return (g.elves[utils.Coordinate{X: c.X - 1, Y: c.Y + 1}] || g.elves[utils.Coordinate{X: c.X - 1, Y: c.Y}] || g.elves[utils.Coordinate{X: c.X - 1, Y: c.Y - 1}])
}

func loadInput(filename string) (g grid) {
	g = grid{elves: map[utils.Coordinate]bool{}}
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			if c == '#' {
				g.addElf(utils.Coordinate{X: i, Y: y})
			}
		}
		y++
	}
	utils.CheckErr(scanner.Err())
	return g
}

func proposeMove(g *grid, start utils.Coordinate, dirn int) utils.Coordinate {
	// if there are no other elves in the 8 coords, do not do anything.
	if !(g.isElfNorth(start) || g.isElfSouth(start) || g.isElfEast(start) || g.isElfWest(start)) {
		return start
	}

	// otherwise if there are other elves in the 8 coords, consider the directions
	for d := dirn; d < dirn+4; d++ {
		switch d % 4 {
		case 0:
			// N
			if !g.isElfNorth(start) {
				return utils.Coordinate{X: start.X, Y: start.Y - 1}
			}
			//fmt.Println("found elf north", start, g)
		case 1:
			// S
			if !g.isElfSouth(start) {
				return utils.Coordinate{X: start.X, Y: start.Y + 1}
			}
		case 2:
			// W
			if !g.isElfWest(start) {
				return utils.Coordinate{X: start.X - 1, Y: start.Y}
			}
		case 3:
			// E
			if !g.isElfEast(start) {
				return utils.Coordinate{X: start.X + 1, Y: start.Y}
			}
		}
	}
	return start
}

func acceptMoves(proposedMoves map[utils.Coordinate][]utils.Coordinate) (g grid) {
	g = grid{elves: map[utils.Coordinate]bool{}}
	// for each proposed move, check if it is valid
	// if it is valid, execute it
	for nextPosition, fromPositions := range proposedMoves {
		if len(fromPositions) > 1 {
			for _, fromPosition := range fromPositions {
				g.addElf(fromPosition)
			}
		} else {
			g.addElf(nextPosition)
		}
	}
	// if it is not valid, do not execute it
	// return the new grid
	return g
}
func part1(filename string) int {
	g := loadInput(filename)
	nRounds := 10
	for i := 0; i <= nRounds; i++ {
		// for every round
		proposedMoves := map[utils.Coordinate][]utils.Coordinate{}
		for position := range g.elves {
			// for each elf, propose a valid move
			// propose a move
			nextPosition := proposeMove(&g, position, i%4)
			proposedMoves[nextPosition] = append(proposedMoves[nextPosition], position)

		}
		// check that moves are still valid given other elves moves, execute if valid
		g = acceptMoves(proposedMoves)
	}

	return g.nEmptyTiles()
}

func part2(filename string) int {
	g := loadInput(filename)
	nRounds := 1000000
	for i := 0; i <= nRounds; i++ {
		// for every round
		proposedMoves := map[utils.Coordinate][]utils.Coordinate{}
		for position := range g.elves {
			// for each elf, propose a valid move
			// propose a move
			nextPosition := proposeMove(&g, position, i%4)
			proposedMoves[nextPosition] = append(proposedMoves[nextPosition], position)

		}
		// check that moves are still valid given other elves moves, execute if valid
		g2 := acceptMoves(proposedMoves)
		if isGridsEqual(&g, &g2) {
			return i + 1
		}
		g = g2
	}

	return g.nEmptyTiles()
}

func isGridsEqual(g1, g2 *grid) bool {
	if len(g1.elves) != len(g2.elves) {
		return false
	}
	for e := range g1.elves {
		if !g2.isElf(e) {
			return false
		}
	}
	return true
}
