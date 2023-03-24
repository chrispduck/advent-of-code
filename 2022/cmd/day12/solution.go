package main

import (
	"advent-of-code/cmd/utils"
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	fmt.Println(part1("example_input.txt"))
	fmt.Println(part1("input.txt"))
	fmt.Println(part2("example_input.txt"))
	fmt.Println(part2("input.txt"))
}

func loadInput(filename string) (grid [][]rune) {
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	utils.CheckErr(scanner.Err())
	return grid
}

type coordinate struct {
	x, y int
}

func findLetter(letter rune, grid [][]rune) coordinate {
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			if grid[j][i] == letter {
				return coordinate{
					x: i,
					y: j,
				}
			}
		}
	}
	log.Fatalln("failed to find letter")
	return coordinate{}
}
func part1(filename string) int {
	grid := loadInput(filename)
	start := findLetter('S', grid)
	finish := findLetter('E', grid)
	// replace with heights
	grid[start.y][start.x] = 'a'
	grid[finish.y][finish.x] = 'z'
	moves, err := findMinMoves(start, finish, grid)
	utils.CheckErr(err)
	return moves
}

func neighbourCoords(c coordinate) []coordinate {
	return []coordinate{
		{
			x: c.x + 1,
			y: c.y,
		},
		{
			x: c.x - 1,
			y: c.y,
		},
		{
			x: c.x,
			y: c.y + 1,
		},
		{
			x: c.x,
			y: c.y - 1,
		},
	}
}
func findMinMoves(start, finish coordinate, grid [][]rune) (int, error) {
	// initialise
	reachableIn := map[coordinate]int{start: 0}
	move := 0
	newCoords := []coordinate{start}
	sizeX := len(grid[0])
	sizeY := len(grid)
	for {
		// while not found the finish, make a move from every new coordinate
		move += 1
		if len(newCoords) == 0 {
			return -1, errors.New("failed to find route")
		}
		var nextCoords []coordinate
		for _, newCoord := range newCoords {
			// try moving in every direction
			possibleNextCoords := neighbourCoords(newCoord)
			for _, c := range possibleNextCoords {
				// is it a valid move?
				inTheGrid := c.x >= 0 && c.y >= 0 && c.x < sizeX && c.y < sizeY
				if !inTheGrid {
					continue
				}
				stepLessThanOne := int(grid[c.y][c.x]-grid[newCoord.y][newCoord.x]) <= 1
				if !stepLessThanOne {
					continue
				}
				_, alreadyReached := reachableIn[c]
				if alreadyReached {
					continue

				}
				// final move
				if c == finish {
					return move, nil
				}
				// add squares reached and to deviate from in the next moves
				reachableIn[c] = move
				nextCoords = append(nextCoords, c)
			}
		}
		newCoords = nextCoords
	}
}

func findAllLetters(letter rune, grid [][]rune) []coordinate {
	var coords []coordinate
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			if grid[j][i] == letter {
				coords = append(coords, coordinate{
					x: i,
					y: j,
				})
			}
		}
	}
	return coords
}
func part2(filename string) int {
	grid := loadInput(filename)
	start := findLetter('S', grid)
	finish := findLetter('E', grid)

	// replace with heights
	grid[start.y][start.x] = 'a'
	grid[finish.y][finish.x] = 'z'

	// for all letter a
	possibleStartingPositions := findAllLetters('a', grid)
	shortestPathLength := math.MaxInt
	for _, start := range possibleStartingPositions {
		// some starting points are isolated from finish, handle err
		pathLength, err := findMinMoves(start, finish, grid)
		if err == nil && pathLength < shortestPathLength {
			shortestPathLength = pathLength
		}
	}
	return shortestPathLength
}
