package main

import (
	"advent-of-code/cmd/utils"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(part1("example_input.txt"))
	fmt.Println(part1("input.txt"))
	fmt.Println(part2("example_input.txt"))
	fmt.Println(part2("input.txt"))
}

type line struct {
	c1 utils.Coordinate
	c2 utils.Coordinate
}

func loadInput(filename string) (lines []line) {
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)
	r := regexp.MustCompile("\\d+")
	for scanner.Scan() {
		text := scanner.Text()
		coordsStr := strings.Split(text, "->")
		size := len(coordsStr)
		for i := 0; i < size-1; i++ {
			c1Arr := r.FindAllString(coordsStr[i], -1)
			c2Arr := r.FindAllString(coordsStr[i+1], -1)
			//fmt.Println(c1Arr, c2Arr)
			lines = append(lines, line{
				c1: utils.Coordinate{
					utils.StrToInt(c1Arr[0]),
					utils.StrToInt(c1Arr[1]),
				},
				c2: utils.Coordinate{
					X: utils.StrToInt(c2Arr[0]),
					Y: utils.StrToInt(c2Arr[1]),
				},
			})
		}

	}
	utils.CheckErr(scanner.Err())
	return lines
}

func createGrid(lines []line) map[utils.Coordinate]bool {
	grid := make(map[utils.Coordinate]bool)
	for _, l := range lines {
		grid = drawLine(l, grid)
	}
	return grid
}

func drawLine(l line, grid map[utils.Coordinate]bool) map[utils.Coordinate]bool {
	dx := l.c2.X - l.c1.X
	dy := l.c2.Y - l.c1.Y

	// horizontal
	if dy == 0 && dx != 0 {
		sign := dx / utils.Abs(dx)
		for i := 0; i <= utils.Abs(dx); i++ {
			// move in the direction of
			grid[utils.Coordinate{
				l.c1.X + sign*i,
				l.c1.Y,
			}] = true
		}
		return grid
	}

	// vertical
	if dx == 0 && dy != 0 {
		sign := dy / utils.Abs(dy)
		for i := 0; i <= utils.Abs(dy); i++ {
			// move in the direction of
			grid[utils.Coordinate{
				X: l.c1.X,
				Y: l.c1.Y + sign*i,
			}] = true
		}
		return grid
	}

	// unknown
	log.Fatalln("unknown line direction", l)
	return grid
}

func placeSand(start utils.Coordinate, yEnd int, grid map[utils.Coordinate]bool) (updatedGrid map[utils.Coordinate]bool, exit bool) {
	sand := start
	down := utils.Coordinate{0, 1}
	downleft := utils.Coordinate{-1, 1}
	downright := utils.Coordinate{1, 1}
	for {
		if sand.Y > yEnd {
			// program termination - free-falling sand
			return updatedGrid, true
		}
		// try to move
		if !occupied(sand.Add(down), grid) {
			sand = sand.Add(down)
			continue
		}
		if !occupied(sand.Add(downleft), grid) {
			sand = sand.Add(downleft)
			continue
		}
		if !occupied(sand.Add(downright), grid) {
			sand = sand.Add(downright)
			continue
		}
		// nowhere to go. place the sand
		grid[sand] = true
		return grid, false
	}
}

func occupied(c utils.Coordinate, grid map[utils.Coordinate]bool) bool {
	_, exists := grid[c]
	return exists
}

func findMaxY(grid map[utils.Coordinate]bool) int {
	maxY := math.MinInt
	for c, _ := range grid {
		if c.Y > maxY {
			maxY = c.Y
		}
	}
	return maxY
}

func part1(filename string) int {
	lines := loadInput(filename)
	grid := createGrid(lines)
	start := utils.Coordinate{500, 0}
	yEnd := findMaxY(grid)
	nGrainsOfSand := 0
	exit := false
	for {
		grid, exit = placeSand(start, yEnd, grid)
		if exit {
			return nGrainsOfSand
		}
		nGrainsOfSand += 1
	}
}

func placeSandPart2(start utils.Coordinate, yFloor int, grid map[utils.Coordinate]bool) (updatedGrid map[utils.Coordinate]bool, exit bool) {
	sand := start
	down := utils.Coordinate{0, 1}
	downleft := utils.Coordinate{-1, 1}
	downright := utils.Coordinate{1, 1}
	for {
		// implement floor logic
		if sand.Y-1 >= yFloor {
			grid[sand] = true
			return grid, false
		}
		// try to move
		if !occupied(sand.Add(down), grid) {
			sand = sand.Add(down)
			continue
		}
		if !occupied(sand.Add(downleft), grid) {
			sand = sand.Add(downleft)
			continue
		}
		if !occupied(sand.Add(downright), grid) {
			sand = sand.Add(downright)
			continue
		}
		// nowhere to go. place the sand
		if sand == start {
			return grid, true
		}
		grid[sand] = true
		return grid, false
	}
}

func part2(filename string) int {
	lines := loadInput(filename)
	grid := createGrid(lines)
	start := utils.Coordinate{500, 0}
	yEnd := findMaxY(grid)
	nGrainsOfSand := 0
	exit := false
	for {
		grid, exit = placeSandPart2(start, yEnd, grid)
		nGrainsOfSand += 1 // last sand is included in count
		if exit {
			return nGrainsOfSand
		}
	}
}
