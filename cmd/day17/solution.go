package main

import (
	"advent-of-code/cmd/utils"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const (
	vertOffset = 3
)

func main() {
	//fmt.Println(part1("example_input.txt", 2022))
	//fmt.Println(part1("input.txt", 2022))
	fmt.Println(part2("example_input.txt", 1_000_000_000_000))
	fmt.Println(part2("input.txt", 1_000_000_000_000))
}

type Shape struct {
	//uLim, lLim, rLim, dLim int
	Coords []utils.Coordinate
}

func (s Shape) move(x utils.Coordinate) Shape {
	var res Shape
	res.Coords = make([]utils.Coordinate, len(s.Coords))
	for idx, _ := range s.Coords {
		res.Coords[idx] = s.Coords[idx].Add(x)
	}
	return res
}

func (s Shape) leftOne() Shape {
	return s.move(utils.Coordinate{-1, 0})
}

func (s Shape) rightOne() Shape {
	return s.move(utils.Coordinate{1, 0})
}

func (s Shape) downOne() Shape {
	return s.move(utils.Coordinate{0, -1})
}

func (s Shape) yMax() int {
	max := math.MinInt64
	for _, c := range s.Coords {
		if c.Y > max {
			max = c.Y
		}
	}
	return max
}
func (s Shape) contains(c utils.Coordinate) bool {
	for _, coord := range s.Coords {
		if coord == c {
			return true
		}
	}
	return false
}

func (s Shape) hitWall() bool {
	for i := 0; i < len(s.Coords); i++ {
		if s.Coords[i].X < 0 || s.Coords[i].X >= 7 {
			return true
		}
	}
	return false
}

func attemptLRMove(move rune, s Shape, grid *[][]bool) Shape {
	var res Shape
	switch move {
	case '<':
		res = s.leftOne()
	case '>':
		res = s.rightOne()
	}
	isHitWall := res.hitWall()
	if isHitWall {
		return s
	}
	isGridCollision := gridCollision(res, grid)
	if isGridCollision {
		return s
	}
	return res
}

func attemptDownMove(s Shape, grid *[][]bool) (isHitBottom bool, s2 Shape) {
	res := s.downOne()
	if isGridCollision := gridCollision(res, grid); isGridCollision {
		return true, s
	}
	return false, res
}

func addToGrid(s Shape, grid *[][]bool) {
	for _, c := range s.Coords {
		(*grid)[c.Y][c.X] = true
	}
}

func gridCollision(s Shape, grid *[][]bool) bool {
	for _, c := range s.Coords {
		if c.Y < 0 || (*grid)[c.Y][c.X] {
			return true
		}
	}
	return false
}

func printGrid(s Shape, grid [][]bool) {
	for y := len(grid) - 1; y >= 0; y-- {
		line := "|"
		for x := 0; x < len(grid[0]); x++ {
			if s.contains(utils.Coordinate{x, y}) {
				line += "@"
			} else if grid[y][x] == true {
				line += "#"
			} else {
				line += "."
			}
		}
		line += "|"
		fmt.Println(line)
	}
	fmt.Println("+-------+")
}

func loadInput(filename string) (cmds []rune) {
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		for _, r := range line {
			cmds = append(cmds, r)
		}
		return cmds
	}
	utils.CheckErr(scanner.Err())
	return []rune{}
}

func createShapes() []Shape {
	horizontal := Shape{
		Coords: []utils.Coordinate{{2, 0}, {3, 0}, {4, 0}, {5, 0}},
	}
	plus := Shape{
		Coords: []utils.Coordinate{{3, 2}, {2, 1}, {3, 1}, {4, 1}, {3, 0}},
	}
	backwardL := Shape{
		Coords: []utils.Coordinate{{2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}},
	}
	vert := Shape{
		Coords: []utils.Coordinate{{2, 0}, {2, 1}, {2, 2}, {2, 3}},
	}
	square := Shape{
		Coords: []utils.Coordinate{{2, 0}, {3, 0}, {2, 1}, {3, 1}},
	}
	allShapes := []Shape{horizontal, plus, backwardL, vert, square}
	return allShapes
}

func boolToString(b bool) string {
	if b {
		return "#"
	}
	return "."
}

// HOW TO FIND A CYCLE
// Have we seen this piece before, this command position, and the top 30 rows of the rock formation the same?
// If so, what position was it?

// idxCmd 0 based, peice 0 based idx
func stateString(cmdIdx int, peiceIdx int, grid [][]bool) string {
	gridString := ""
	if len(grid) != 30 {
		log.Fatalln("grid is :", len(grid))
	}
	for y := 0; y < 30; y++ {
		for x := 0; x < 7; x++ {
			gridString += boolToString(grid[y][x])
		}
	}
	return strconv.Itoa(cmdIdx) + "-" + strconv.Itoa(peiceIdx) + "" + gridString
}

func getLastNRows(n, ymax int, grid [][]bool) [][]bool {
	if ymax < n {
		return grid[:n]
	}
	return grid[ymax-n : ymax]
}

func part1(filename string, nRocks int) int {
	cmds := loadInput(filename)
	allShapes := createShapes()
	nShapes := len(allShapes)
	ymax := -1
	idxCmd := 0
	m, n := utils.Min(nRocks*4, 1_000_000), 7
	grid := make([][]bool, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]bool, n)
	}

	// to find cycles & implement skip
	heights := make(map[string]int)
	rocksPlaced := make(map[string]int)
	toAddheight := 0
	foundLoop := false

	for i := 0; i < nRocks; i++ {
		// check for cycles
		upperGrid := getLastNRows(30, ymax, grid)
		stateString := stateString(idxCmd, i%nShapes, upperGrid)
		yPrev, exists := heights[stateString]
		if exists && !foundLoop {
			foundLoop = true
			loopLength := i - rocksPlaced[stateString]
			loopHeight := ymax - yPrev
			rocksRemaining := nRocks - i
			repetitions := rocksRemaining / loopLength

			// Add the height at the end
			toAddheight = repetitions * loopHeight
			i += repetitions * loopLength
		}
		heights[stateString] = ymax
		rocksPlaced[stateString] = i

		// Place shape
		// create the right new shape
		shapeToPlace := allShapes[i%nShapes]

		// put it in the correct v offset compared with the grid (wrt to the yMax of the grid)
		shapeToPlace = shapeToPlace.move(utils.Coordinate{X: 0, Y: vertOffset + ymax + 1})
		for {
			// move it across if possible
			shapeToPlace = attemptLRMove(cmds[idxCmd], shapeToPlace, &grid)
			idxCmd++
			idxCmd = idxCmd % len(cmds) // never reuse
			// move down if possible, and repeat else break
			var isHitBottom bool
			isHitBottom, shapeToPlace = attemptDownMove(shapeToPlace, &grid)

			// place shape into grid, and update yMax
			if isHitBottom {
				addToGrid(shapeToPlace, &grid)
				if shapeToPlace.yMax() > ymax {
					ymax = shapeToPlace.yMax()
				}
				break
			}
		}
	}
	return ymax + 1 + toAddheight
}

func part2(filename string, nRocks int) int {
	return part1(filename, nRocks)
}
