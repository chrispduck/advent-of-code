package main

import (
	"advent-of-code/cmd/utils"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(part1("example_input.txt"))
	fmt.Println(part1("input.txt"))
	//fmt.Println(part2("example_input.txt"))
	//fmt.Println(part2("input.txt"))
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

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Printf("%c\n", row)
	}
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
	printGrid(grid)
	start := findLetter('S', grid)
	finish := findLetter('E', grid)
	// replace with heights
	grid[start.y][start.x] = 'a'
	grid[finish.y][finish.x] = 'z'
	sizeX := len(grid[0])
	sizeY := len(grid)

	// initialise
	reachableIn := map[coordinate]int{start: 0}
	move := 0
	newCoords := []coordinate{start}

	for {
		// while not found the finish, make a move from every new coordinate
		move += 1
		fmt.Println(move, newCoords)
		if len(newCoords) == 0 {
			log.Fatalln("got stuck")
		}
		var nextCoords []coordinate
		for _, newCoord := range newCoords {
			// try moving in every direction
			possibleNextCoords := []coordinate{
				{
					x: newCoord.x + 1,
					y: newCoord.y,
				},
				{
					x: newCoord.x - 1,
					y: newCoord.y,
				},
				{
					x: newCoord.x,
					y: newCoord.y + 1,
				},
				{
					x: newCoord.x,
					y: newCoord.y - 1,
				},
			}
			for _, c := range possibleNextCoords {

				// is it a valid move?
				inTheGrid := c.x >= 0 && c.y >= 0 && c.x < sizeX && c.y < sizeY
				if !inTheGrid {
					continue
				}
				stepLessThanOne := utils.Abs(int(grid[newCoord.y][newCoord.x]-grid[c.y][c.x])) <= 1
				if !stepLessThanOne {
					continue
				}
				_, alreadyReached := reachableIn[c]
				if alreadyReached {
					continue

				}
				// final move
				if c == finish {
					return move
				}
				// add squares reached and to deviate from in the next moves
				reachableIn[c] = move
				nextCoords = append(nextCoords, c)
			}

		}
		newCoords = nextCoords
	}
}

func part2(filename string) int {
	return 0
}
