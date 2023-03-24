package main

import (
	"advent-of-code/cmd/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// >> me
func main() {
	fmt.Println(part1("example_input.txt"))
	fmt.Println(part1("input.txt"))
	//fmt.Println(part2("example_input.txt"))
	//fmt.Println(part2("input.txt"))
}

// 0 = north (up)
// 1 = east (right)
// 2 = south (down)
// 3 = west (left)
//direction := 0

func rotateLeft(direction int) int {
	return (direction + 3) % 4
}

func rotateRight(direction int) int {
	return (direction + 1) % 4
}

type grid struct {
	grid [][]rune
}

func (g *grid) print() {
	for _, row := range g.grid {
		for _, c := range row {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

func (g *grid) isEmpty(coordinate utils.Coordinate) bool {
	return g.grid[coordinate.Y][coordinate.X] == ' '
}

func (g *grid) isWall(coordinate utils.Coordinate) bool {
	return g.grid[coordinate.Y][coordinate.X] == '#'
}

func (g *grid) isValid(coordinate utils.Coordinate) bool {
	return g.grid[coordinate.Y][coordinate.X] == '.'
}

func (g *grid) moveOne(direction int, position utils.Coordinate) utils.Coordinate {
	nextPosition := position
	switch direction {
	case 0:
		nextPosition.Y--
	case 1:
		nextPosition.X++
	case 2:
		nextPosition.Y++
	case 3:
		nextPosition.X--
	}
	nextPosition = nextPosition.WrapAround(len(g.grid[0]), len(g.grid))
	if g.isValid(nextPosition) {
		return nextPosition
	} else if g.isEmpty(nextPosition) {
		attemptedPosition := g.moveOne(direction, nextPosition)
		if !g.isValid(attemptedPosition) {
			return position
		} else {
			return attemptedPosition
		}
	} else if g.isWall(nextPosition) {
		return position
	}
	return position
}

func loadInput(filename string) (g *grid, nums []int, directions []string) {
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)
	y := 0
	g = &grid{}
	grid_parsed := false
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		if line == "" {
			grid_parsed = true
			continue
		}
		if !grid_parsed {
			row := []rune(line)
			if l := len(g.grid); l > 0 {
				nPaddingSpaces := len(g.grid[0]) - len(row)
				//fmt.Println("padding row", l, len(line))
				for i := 0; i < nPaddingSpaces; i++ {
					//fmt.Println("padding row once", i)
					row = append(row, ' ')
				}
			}
			g.grid = append(g.grid, row)
			y++
		} else {
			r := regexp.MustCompile("\\d+")
			numsStrs := r.FindAllString(line, -1)
			for _, numStr := range numsStrs {
				nums = append(nums, utils.StrToInt(numStr))
			}
			r = regexp.MustCompile("[LR]")
			directions = r.FindAllString(line, -1)

		}
	}
	utils.CheckErr(scanner.Err())

	return
}

func computePassword(position utils.Coordinate, direction int) int {
	//fmt.Println(direction)
	scoringDirection := (direction + 3) % 4
	//fmt.Println("scoringDirection", scoringDirection)
	return 1000*(position.Y+1) + 4*(position.X+1) + scoringDirection
}

func part1(filename string) int {
	g, nums, directions := loadInput(filename)
	//fmt.Println(nums, directions)
	if len(nums) != len(directions)+1 {
		panic("bad input")
	}
	//fmt.Printf("%+v", g)
	//g.print()
	position := utils.Coordinate{X: 0, Y: 0}
	direction := 1 // starts pointing right
	for i := 0; i < len(nums); i++ {
		if i != 0 {
			//fmt.Println("rotating", directions[i-1])
			switch directions[i-1] {
			case "L":
				direction = rotateLeft(direction)
			case "R":
				direction = rotateRight(direction)
			}
		}
		//fmt.Println("moving by ", nums[i])
		for j := 0; j < nums[i]; j++ {
			position = g.moveOne(direction, position)
			//fmt.Printf("moved to %+v", position)
		}
	}
	return computePassword(position, direction)
}

//func part2(filename string) int {
//	return 0
//}
